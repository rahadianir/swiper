package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/rahadianir/swiper/internal/auth"
	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/config"
	"github.com/rahadianir/swiper/internal/pkg/cache"
	mw "github.com/rahadianir/swiper/internal/pkg/middleware"
	"github.com/rahadianir/swiper/internal/swiper"
	"github.com/rahadianir/swiper/internal/users"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

func SetupDependencies(ctx context.Context) *common.Dependencies {
	cfg := config.LoadConfig()

	// setup database
	db, err := sqlx.ConnectContext(ctx, "postgres", cfg.DatabaseURI)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if redisClient == nil {
		log.Fatal("failed to init redis client")
	}

	validator := validator.New()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	return &common.Dependencies{
		Config:      cfg,
		DB:          db,
		RedisClient: redisClient,
		Validator:   validator,
	}
}

func InitRoutes(deps *common.Dependencies) http.Handler {
	// wiring shared package layer
	cacheStore := cache.NewCacheStore(deps)

	// wiring up repository layer
	userRepo := users.NewUserRepo(deps)
	swiperRepo := swiper.NewSwipeRepo(deps)

	// wiring up logic layer
	authLogic := auth.NewAuthLogic(deps)
	userLogic := users.NewUserLogic(deps, userRepo, authLogic)
	swiperLogic := swiper.NewSwiperLogic(deps, userRepo, cacheStore, swiperRepo)

	// wiring up handler layer
	userHandler := users.NewUserHandler(deps, userLogic)
	swiperHandler := swiper.NewSwiperHandler(deps, swiperLogic)

	// init custom middlewares
	authMiddleware := mw.AuthMiddleware{
		Dependencies: deps,
	}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test!"))
	})

	r.Post("/signup", userHandler.Register)
	r.Post("/signin", userHandler.Login)
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.ValidateToken)
		r.Get("/profile/{id}", userHandler.GetProfileByID)

		r.Get("/queue", swiperHandler.GetTargetProfile)
		r.Post("/swipe/right/{id}", swiperHandler.SwipeRight)
		r.Post("/swipe/left/{id}", swiperHandler.SwipeLeft)
	})

	return r
}

func StartServer(ctx context.Context) {
	// Server run context
	serverCtx, serverStop := context.WithCancel(ctx)

	deps := SetupDependencies(serverCtx)
	r := InitRoutes(deps)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", deps.Config.HTTPPort),
		Handler: r,
	}

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		err = deps.DB.Close()
		if err != nil {
			log.Fatal(err)
		}

		serverStop()
	}()

	// Run the server
	log.Println("server starts!")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

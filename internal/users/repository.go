package users

import (
	"context"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
)

type UserRepo struct {
	*common.Dependencies
}

func NewUserRepo(deps *common.Dependencies) *UserRepo {
	return &UserRepo{
		Dependencies: deps,
	}
}

func (r *UserRepo) Register(ctx context.Context, user models.User) (int, error) {
	var userID int
	q := `INSERT INTO users (name, username, password, age, gender, location, is_premium, is_verified, created_at) VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, now()) RETURNING id`

	err := r.DB.QueryRowContext(ctx, q, user.Name, user.Username, user.Password, user.Age, user.Gender, user.Location, user.IsPremium, user.IsVerified).Scan(&userID)
	if err != nil {
		// parse error
		if strings.Contains(err.Error(), "unique") {
			return userID, fmt.Errorf("duplicate username")
		}
		return userID, err
	}

	return userID, nil
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	q := `SELECT
			id, 
			name, 
			username, 
			password, 
			age, 
			gender, 
			location, 
			is_premium AS ispremium, 
			is_verified AS isverified, 
			created_at AS createdat, 
			updated_at AS updatedat, 
			deleted_at AS deletedat
		FROM
			users u
		WHERE 
			u.username = $1
			AND
			u.deleted_at ISNULL;`

	err := r.DB.QueryRowxContext(ctx, q, username).StructScan(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByUserID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	q := `SELECT
			id, 
			name, 
			username, 
			password, 
			age, 
			gender, 
			location, 
			is_premium AS ispremium, 
			is_verified AS isverified, 
			created_at AS createdat, 
			updated_at AS updatedat, 
			deleted_at AS deletedat
		FROM
			users u
		WHERE 
			u.id = $1
			AND
			u.deleted_at ISNULL;`

	err := r.DB.QueryRowxContext(ctx, q, id).StructScan(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetRandomUser(ctx context.Context, excludeList []int) (models.User, error) {
	var user models.User
	q := `SELECT
			id, 
			name, 
			username, 
			password, 
			age, 
			gender, 
			location, 
			is_premium AS ispremium, 
			is_verified AS isverified, 
			created_at AS createdat, 
			updated_at AS updatedat, 
			deleted_at AS deletedat
		FROM
			users u
		WHERE 
			u.id <> ALL($1)
			AND
			u.deleted_at ISNULL
		ORDER BY 
			RANDOM()
		LIMIT 
			1;`

	err := r.DB.QueryRowxContext(ctx, q, pq.Array(excludeList)).StructScan(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

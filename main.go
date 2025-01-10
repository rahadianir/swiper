package main

import (
	"context"

	"github.com/rahadianir/swiper/internal/server"
)

func main() {
	server.StartServer(context.Background())
}

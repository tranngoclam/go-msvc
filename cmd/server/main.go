package main

import (
	"context"
	"github.com/tranngoclam/go-msvc/internal/config"
	"github.com/tranngoclam/go-msvc/internal/handler"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	srv, err := handler.NewServer(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	srv.Serve(ctx)
}

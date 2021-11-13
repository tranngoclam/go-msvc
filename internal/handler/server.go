package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/tranngoclam/go-msvc/internal/config"
	"github.com/tranngoclam/go-msvc/internal/database"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config  config.Config
	handler *Handler
}

func NewServer(
	ctx context.Context,
	cfg config.Config,
) (*Server, error) {
	http.Lis

	db, err := database.New(ctx, cfg.Database())
	if err != nil {
		return nil, err
	}

	h := NewHandler(&database.Task{DB: db})

	s := &Server{
		config:  cfg,
		handler: h,
	}

	return s, nil
}

func (s *Server) Serve(ctx context.Context) error {
	srv := &http.Server{
		Handler: s.handler.mux,
	}

	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		errCh <- srv.Shutdown(shutdownCtx)
	}()

	http.ListenAndServe(s.config.)

	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	if err := <-errCh; err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	return nil
}

package server

import (
	"backend-assignment/config"
	"backend-assignment/database"
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Config config.Config
	server *http.Server
	db     database.DB
}

func Init(ctx context.Context, db database.DB) (*Server, error) {
	conf := config.Get()

	server := &Server{
		Config: conf,
		db:     db,
	}

	server.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", server.Config.HTTP.Port),
		Handler: router(ctx, server).Handler(),
	}

	return server, nil
}
func (s *Server) Start(ctx context.Context) error {
	// Run communicator in a separate goroutine
	// go communicator.RunCommunicator(ctx)

	go s.gracefulHandler(ctx)
	err := s.server.ListenAndServe()
	return errors.Wrap(err, "server.Start")
}
func (s *Server) gracefulHandler(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Info().Msg("shutting down server")
		s.Shutdown(context.Background())
	}
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

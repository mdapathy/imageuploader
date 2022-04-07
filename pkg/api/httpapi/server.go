package httpapi

import (
	"context"
	"errors"
	"github.com/mdapathy/imageuploader/pkg/config"
	"github.com/mdapathy/imageuploader/pkg/domain"
	"github.com/mdapathy/imageuploader/pkg/domain/query"
	"net/http"
	"time"
)

type resources struct {
	images  domain.ImageService
	queries *query.Factory
}

type Server struct {
	*http.Server
	shutdownTimeout time.Duration
}

func NewServer(addr string, conf *config.Server, image domain.ImageService, factory *query.Factory) *Server {
	resources := resources{
		images:  image,
		queries: factory,
	}

	return &Server{
		Server: &http.Server{
			Addr:           addr,
			Handler:        configureRoutes(&resources),
			ReadTimeout:    conf.ReadTimeout.Duration,
			WriteTimeout:   100 * time.Second,
			IdleTimeout:    conf.IdleTimeout.Duration,
			MaxHeaderBytes: 1 << 20,
		},
		shutdownTimeout: conf.ShutdownTimeout.Duration,
	}
}

func (s *Server) Start(ctx context.Context) error {
	errs := make(chan error)
	go func() {
		errs <- s.ListenAndServe()
	}()

	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		err := s.Shutdown(ctxShutDown)
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}
}

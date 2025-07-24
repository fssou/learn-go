package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type AdapterConfig struct {
	Host string
	Port string
}

type Adapter struct {
	config *AdapterConfig
	server *http.Server
}

func NewAdapter(handler http.Handler, config *AdapterConfig) *Adapter {
	return &Adapter{
		config: config,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
			Handler: handler,
		},
	}
}

func (s *Adapter) Close() error {
	log.Println("Closing http adapter...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if s.server != nil {
		if err := s.server.Shutdown(ctx); err != nil {
			return errors.New("failed to close server: " + err.Error())
		}
	}
	s.server = nil
	return nil
}

func (s *Adapter) Listen() error {
	err := s.server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Adapter closed")
			return nil
		}
		return errors.New("failed to start server: " + err.Error())
	}
	return err
}

package gateway

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Host         string
	Port         string
	Handler      http.Handler
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Timeout      time.Duration
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port string) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithHandler(h http.Handler) Option {
	return func(s *Server) {

	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = timeout
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.IdleTimeout = timeout
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func (s *Server) Connect() {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      s.Handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		fmt.Printf("Starting server on %s with shutdown timeout %v\n", addr, s.Timeout)
		serverErrors <- server.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
			panic(err)
		}
	case <-stop:
		fmt.Println("Received shutdown signal, initiating graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Server shutdown error: %v\n", err)
			if err := server.Close(); err != nil {
				fmt.Printf("Server close error: %v\n", err)
			}
			panic(err)
		}
		fmt.Println("Server gracefully shut down")
	}
}

func NewServer(opts ...Option) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

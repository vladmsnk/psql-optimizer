package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Host string
	Port int
}

type Server struct {
	GrpcServer *grpc.Server
	addr       string
	lis        net.Listener
}

func New(cfg *Config, opts ...grpc.ServerOption) (*Server, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	return &Server{
		GrpcServer: grpcServer,
		addr:       addr,
		lis:        lis,
	}, nil
}

func (s *Server) Serve() error {
	log.Printf("gRPC server starting on %s", s.addr)

	if err := s.GrpcServer.Serve(s.lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func (s *Server) GracefulStop() {
	log.Printf("gRPC server on %s is shutting down", s.addr)

	s.GrpcServer.GracefulStop()
}

func (s *Server) RunWithGracefulShutdown() {
	errCh := make(chan error, 1)

	go func() {
		if err := s.Serve(); err != nil {
			errCh <- err
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		log.Printf("gRPC server error: %v", err)
	case sig := <-sigCh:
		log.Printf("Received signal: %v. Shutting down...", sig)
	}

	s.GracefulStop()
}

func (s *Server) ServeWithContext(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Serve()
	}()

	select {
	case <-ctx.Done():
		s.GracefulStop()
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}

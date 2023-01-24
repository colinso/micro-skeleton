package api

import (
	"Personal/micro-skeleton/internal/api/grpc"
	"Personal/micro-skeleton/internal/config"
	"fmt"
	"net"
	"net/http"

	ggrpc "google.golang.org/grpc"

	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg        *config.Config
	router     *chi.Mux
	grpcServer *grpc.Server
}

func NewServer(cfg *config.Config, router *chi.Mux, grpcServer *grpc.Server) *Server {
	return &Server{
		cfg:        cfg,
		router:     router,
		grpcServer: grpcServer,
	}
}

func (s Server) Start() error {
	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	log.Infof("Starting up at %s", address)
	return http.ListenAndServe(address, s.router)
}

func (s Server) StartGrpc() error {
	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := ggrpc.NewServer()
	grpc.RegisterMicroSkeletonServer(server, *s.grpcServer)
	log.Infof("Starting gRPC server at %s", address)
	return server.Serve(lis)
}

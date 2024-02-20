package gapi

import (
	"context"
	"net"
	"net/http"
	"sharecycle/configs"
	"sharecycle/foundation/database"
	"sharecycle/internal/pb"
	"sharecycle/pkg/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

// Server serves gRPC request for the app
type Server struct {
	pb.UnimplementedSharecycleServer
	conf *configs.Config
	db   *database.DBV1
	l    logger.Logger
	h    *GrpcHandler
}

// Server creates a new gRPC server
func NewGrpcServer(conf *configs.Config, dbV1 *database.DBV1, l logger.Logger, h *GrpcHandler) *Server {
	return &Server{
		conf: conf,
		db:   dbV1,
		l:    l,
		h:    h,
	}
}

func (s *Server) RunGrpcServer() {

	listener, err := net.Listen("tcp", s.conf.GrpcAPIs.Address)
	if err != nil {
		s.l.Fatal("Cannot create gRPC Listener ", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSharecycleServer(grpcServer, s)
	reflection.Register(grpcServer)

	s.l.Infof("Starting GRPC server on: %s", s.conf.GrpcAPIs.Address)
	err = grpcServer.Serve(listener)
	if err != nil {
		s.l.Fatal("Cannot start gRPC server", err)
	}
}

func (s *Server) RunGatewayServer() {

	jsonOption := runtime.WithMarshalerOption("application/json+pretty", &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			Multiline: true, // Optional, implied by presence of "Indent".
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := pb.RegisterSharecycleHandlerServer(ctx, grpcMux, s)
	if err != nil {
		s.l.Fatal("Cannot register gateway handler server", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", s.conf.APIs.Address)
	if err != nil {
		s.l.Fatal("Cannot create gateway Listener ", err)
	}

	s.l.Infof("Starting HTTP server on: %s", s.conf.APIs.Address)

	err = http.Serve(listener, mux)

	if err != nil {
		s.l.Fatal("Cannot start HTTP server", err)
	}
}

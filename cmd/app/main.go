package main

import "sharecycle/internal/gapi"

func main() {
	// Run GRPC and Gateway server
	s := gapi.GrpcReady()
	go s.RunGatewayServer()
	s.RunGrpcServer()
}

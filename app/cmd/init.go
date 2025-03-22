package main

import (
	desc "psql-optimizer/pkg/optimizer"
	grpclib "psql-optimizer/utils/grpc_server"
)

func runGRPCServer(transport desc.OptimizerAPIServer, grpcServer *grpclib.Server) (*grpclib.Server, error) {
	desc.RegisterOptimizerAPIServer(grpcServer.GrpcServer, transport)
	grpcServer.RunWithGracefulShutdown()

	return grpcServer, nil
}

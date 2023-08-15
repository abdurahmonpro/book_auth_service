package grpc

import (
	"auth_service/config"
	"auth_service/genproto/auth_service"
	"auth_service/grpc/client"
	"auth_service/grpc/service"
	"auth_service/pkg/logger"
	"auth_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	auth_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg, srvc))
	auth_service.RegisterAuthServiceServer(grpcServer, service.NewAuthService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}

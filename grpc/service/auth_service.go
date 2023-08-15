package service

import (
	"auth_service/config"
	"auth_service/genproto/auth_service"
	"auth_service/grpc/client"
	"auth_service/pkg/helper/security"
	"auth_service/pkg/logger"
	"auth_service/storage"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	auth_service.UnimplementedAuthServiceServer
}

func NewAuthService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *authService {
	return &authService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *authService) Login(ctx context.Context, req *auth_service.LoginRequest) (*auth_service.TokenResponse, error) {
	log.Println("Login...")

	errAuth := errors.New("username or password wrong")
	errNotfound := errors.New("password is wrong")

	user, err := s.strg.User().GetUserByUsername(ctx, req.Name)
	if err != nil {
		log.Println(err.Error())
		return nil, status.Errorf(codes.Unauthenticated, errAuth.Error())
	}

	match, err := security.ComparePassword(user.Secret, req.Secret)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "util.ComparePassword: %s", errNotfound.Error())
	}

	if !match {
		return nil, status.Errorf(codes.Unauthenticated, errAuth.Error())
	}

	m := map[string]interface{}{
		"user_id":  user.Id,
		"username": user.Secret,
	}

	tokenStr, err := security.GenerateJWT(m, 10*time.Minute, s.cfg.SecretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "util.GenerateJWT: %s", err.Error())
	}
	fmt.Println(tokenStr)

	return &auth_service.TokenResponse{
		Token: tokenStr,
	}, nil
}



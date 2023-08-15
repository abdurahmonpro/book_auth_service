package service

import (
	"auth_service/config"
	"auth_service/genproto/auth_service"
	"auth_service/grpc/client"
	"auth_service/pkg/helper/security"
	"auth_service/pkg/logger"
	"auth_service/storage"
	"errors"
	"fmt"
	"regexp"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	auth_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *UserService {
	return &UserService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *UserService) Create(ctx context.Context, req *auth_service.CreateUser) (*auth_service.User, error) {
	fmt.Println("----------------------------------------------------------------")
	s.log.Info("---CreateUser--->", logger.Any("req", req))

	if len(req.Secret) < 6 {
		err := fmt.Errorf("password must not be less than 6 characters")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	hashedPassword, err := security.HashPassword(req.Secret)
	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Secret = hashedPassword
	fmt.Println(hashedPassword)

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	email := emailRegex.MatchString(req.Email)
	if !email {
		err = fmt.Errorf("email is not valid")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	pKey, err := s.strg.User().Create(ctx, req)

	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.User().GetByPKey(ctx, pKey)
}

func (s *UserService) CheckUser(ctx context.Context, req *auth_service.CheckUserRequest) (*auth_service.CheckUserResponse, error) {
	s.log.Info("---CheckUser--->", logger.Any("req", req))

	if len(req.Secret) < 6 {
		err := errors.New("invalid password")
		s.log.Error("!!!Login 2--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.User().GetUserByUsername(ctx, req.Name)
	if err != nil {
		s.log.Error("!!!Login 3--->", logger.Error(err))
		return &auth_service.CheckUserResponse{Exists: false, Registered: false}, nil
	}

	if user.Name == req.Name && user.Secret == req.Secret {
		return &auth_service.CheckUserResponse{Exists: true, Registered: true}, nil
	}

	return &auth_service.CheckUserResponse{Exists: true}, nil
}

func (i *UserService) GetByID(ctx context.Context, req *auth_service.UserPK) (resp *auth_service.User, err error) {

	i.log.Info("---GetUserByID------>", logger.Any("req", req))

	resp, err = i.strg.User().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetUserByID->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *UserService) GetList(ctx context.Context, req *auth_service.UserListRequest) (resp *auth_service.UserListResponse, err error) {

	i.log.Info("---GetUsers------>", logger.Any("req", req))

	resp, err = i.strg.User().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetUsers->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

package service

import (
	"auth_service/config"
	"auth_service/genproto/auth_service"
	"auth_service/grpc/client"
	"auth_service/pkg/logger"
	"auth_service/storage"

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
func (s *UserService) Create(ctx context.Context, req *auth_service.CreateUser) (*auth_service.CreateUserResponse, error) {
	s.log.Info("---CreateUser--->", logger.Any("req", req))

	// if len(req.Secret) < 6 {
	// 	err := fmt.Errorf("password must be at least 6 characters")
	// 	s.log.Error("!!!CreateUser--->", logger.Error(err))
	// 	return nil, err
	// }

	// hashedPassword, err := security.HashPassword(req.Secret)
	// if err != nil {
	// 	s.log.Error("!!!CreateUser--->", logger.Error(err))
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }
	// req.Secret = hashedPassword

	// emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	// if !emailRegex.MatchString(req.Email) {
	// 	err = fmt.Errorf("email is not valid")
	// 	s.log.Error("!!!CreateUser--->", logger.Error(err))
	// 	return nil, err
	// }

	pKey, err := s.strg.User().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	createdUser, err := s.strg.User().GetByPKey(ctx, pKey)
	if err != nil {
		s.log.Error("!!!GetByidUser--->", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to fetch created user")
	}

	response := &auth_service.CreateUserResponse{
		Data:    []*auth_service.User{createdUser},
		IsOk:    true,
		Message: "ok",
	}

	return response, nil
}

func (s *UserService) CheckUser(ctx context.Context, req *auth_service.CheckUserRequest) (*auth_service.CheckUserResponse, error) {
	s.log.Info("---CheckUser--->", logger.Any("req", req))

	// if len(req.Secret) < 6 {
	// 	err := errors.New("invalid password")
	// 	s.log.Error("!!!Login 2--->", logger.Error(err))
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }

	user, err := s.strg.User().GetUserByUsername(ctx, &auth_service.GetByName{Name: req.Name})
	if err != nil {
		s.log.Error("!!!Login 3--->", logger.Error(err))
		return &auth_service.CheckUserResponse{Exists: false, Registered: false}, nil
	}

	if user.Name == req.Name && user.Secret == req.Secret {
		return &auth_service.CheckUserResponse{Exists: true, Registered: true}, nil
	}

	return &auth_service.CheckUserResponse{Exists: false, Registered: false}, nil
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

func (i *UserService) GetUserByName(ctx context.Context, req *auth_service.GetByName) (*auth_service.CreateUserResponse, error) {

	i.log.Info("---GetUserByName------>", logger.Any("req", req))

	user, err := i.strg.User().GetUserByUsername(ctx, req)
	if err != nil {
		i.log.Error("!!!GetUserByName->User->Get--->", logger.Error(err))

		response := &auth_service.CreateUserResponse{
			Data:    []*auth_service.User{},
			IsOk:    true,
			Message: "unable to authorize",
		}
		return response, nil
	}

	response := &auth_service.CreateUserResponse{
		Data:    []*auth_service.User{user},
		IsOk:    true,
		Message: "ok",
	}
	return response, nil
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

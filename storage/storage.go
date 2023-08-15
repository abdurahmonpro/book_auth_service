package storage

import (
	"auth_service/genproto/auth_service"

	"context"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(context.Context, *auth_service.CreateUser) (*auth_service.UserPK, error)
	GetByPKey(context.Context, *auth_service.UserPK) (*auth_service.User, error)
	GetAll(context.Context, *auth_service.UserListRequest) (*auth_service.UserListResponse, error)
	GetUserByUsername(context.Context, string)(*auth_service.User, error)
}

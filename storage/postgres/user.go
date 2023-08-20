package postgres

import (
	"auth_service/genproto/auth_service"
	"auth_service/pkg/helper"

	"context"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *auth_service.CreateUser) (resp *auth_service.UserPK, err error) {
	query := `
		INSERT INTO "user" (
			name,
			email,
			key,
			secret,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id
	`

	var id int
	err = u.db.QueryRow(
		ctx,
		query,
		req.Name,
		req.Email,
		req.Key,
		req.Secret,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &auth_service.UserPK{Id: int32(id)}, nil
}

func (u *userRepo) GetByPKey(ctx context.Context, req *auth_service.UserPK) (user *auth_service.User, err error) {

	query := `
		SELECT 
			id,
			name,
			email,
			key,
			secret
		FROM "user"
		WHERE id = $1
	`

	var (
		id     sql.NullInt32
		name   sql.NullString
		email  sql.NullString
		key    sql.NullString
		secret sql.NullString
	)

	err = u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&email,
		&key,
		&secret,
	)
	if err != nil {
		return user, err
	}

	user = &auth_service.User{
		Id:     id.Int32,
		Name:   name.String,
		Email:  email.String,
		Key:    key.String,
		Secret: secret.String,
	}

	return
}

func (u *userRepo) GetAll(ctx context.Context, req *auth_service.UserListRequest) (resp *auth_service.UserListResponse, err error) {
	resp = &auth_service.UserListResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE "
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			email,
			key,
			secret
		FROM "user"
	`
	if len(req.GetSearch()) > 0 {
		filter += " AND (name || ' ' || email) ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}
	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}
	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id     sql.NullInt32
			name   sql.NullString
			email  sql.NullString
			key    sql.NullString
			secret sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&email,
			&key,
			&secret,
		)
		if err != nil {
			return resp, err
		}

		resp.Users = append(resp.Users, &auth_service.User{
			Id:     id.Int32,
			Name:   name.String,
			Email:  email.String,
			Key:    key.String,
			Secret: secret.String,
		})
	}

	return
}

func (u *userRepo) GetUserByUsername(ctx context.Context, req *auth_service.GetByName) (resp *auth_service.User, err error) {
	query := `
		SELECT 
			id,
			name,
			email,
			key,
			secret
		FROM "user"
		WHERE name = $1
	`

	var (
		id     sql.NullInt32
		name   sql.NullString
		email  sql.NullString
		key    sql.NullString
		secret sql.NullString
	)

	err = u.db.QueryRow(ctx, query, req.Name).Scan(
		&id,
		&name,
		&email,
		&key,
		&secret,
	)
	if err != nil {
		return resp, err
	}

	resp = &auth_service.User{
		Id:     id.Int32,
		Name:   name.String,
		Email:  email.String,
		Key:    key.String,
		Secret: secret.String,
	}

	return
}

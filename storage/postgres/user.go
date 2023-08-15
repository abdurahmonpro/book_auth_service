package postgres

import (
	"auth_service/genproto/auth_service"
	"auth_service/pkg/helper"

	"context"
	"database/sql"

	"github.com/google/uuid"
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

	id := uuid.New().String()

	query := `
		INSERT INTO "user" (
			id,
			name,
			email,
			key,
			secret,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`

	_, err = u.db.Exec(
		ctx,
		query,
		id,
		req.Name,
		req.Email,
		req.Key,
		req.Secret,
	)
	if err != nil {
		return nil, err
	}

	return &auth_service.UserPK{Id: id}, nil
}

func (u *userRepo) GetByPKey(ctx context.Context, req *auth_service.UserPK) (user *auth_service.User, err error) {

	query := `
		SELECT 
			id,
			name,
			email,
			key,
			secret,
			created_at,
			updated_at
		FROM "user"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		name       sql.NullString
		email      sql.NullString
		key        sql.NullString
		secret     sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	err = u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&email,
		&key,
		&secret,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return user, err
	}

	user = &auth_service.User{
		Id:        id.String,
		Name:      name.String,
		Email:     email.String,
		Key:       key.String,
		Secret:    secret.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
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
			secret,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
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
			id         sql.NullString
			name       sql.NullString
			email      sql.NullString
			key        sql.NullString
			secret     sql.NullString
			created_at sql.NullString
			updated_at sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&email,
			&key,
			&secret,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return resp, err
		}

		resp.Users = append(resp.Users, &auth_service.User{
			Id:        id.String,
			Name:      name.String,
			Email:     email.String,
			Key:       key.String,
			Secret:    secret.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}

	return
}

func (u *userRepo) GetUserByUsername(ctx context.Context, username string)(resp *auth_service.User, err error){
	query := `
		SELECT 
			id,
			name,
			email,
			key,
			secret,
			created_at,
			updated_at
		FROM "user"
		WHERE name = $1
	`

	var (
		id         sql.NullString
		name       sql.NullString
		email      sql.NullString
		key        sql.NullString
		secret     sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	err = u.db.QueryRow(ctx, query, username).Scan(
		&id,
		&name,
		&email,
		&key,
		&secret,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return resp, err
	}

	resp = &auth_service.User{
		Id:        id.String,
		Name:      name.String,
		Email:     email.String,
		Key:       key.String,
		Secret:    secret.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}

	return
}

// func (u *userRepo) Update(ctx context.Context, req *auth_service.Update) (rowsAffected int64, err error) {

// 	var (
// 		query  string
// 		params map[string]interface{}
// 	)

// 	query = `
// 		UPDATE
// 			"user"
// 		SET
// 			name = :name,
// 			email = :email,
// 			key = :key,
// 			secret = :secret,
// 			updated_at = now()
// 		WHERE id = :id
// 	`
// 	params = map[string]interface{}{
// 		"id":     req.GetId(),
// 		"name":   req.GetLastName(),
// 		"email":  req.GetLastName(),
// 		"key":    req.GetPhoneNumber(),
// 		"secret": req.GetDateOfBirth(),
// 	}

// 	query, args := helper.ReplaceQueryParams(query, params)

// 	result, err := u.db.Exec(ctx, query, args...)
// 	if err != nil {
// 		return
// 	}

// 	return result.RowsAffected(), nil
// }

// func (u *userRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

// 	var (
// 		set   = " SET "
// 		ind   = 0
// 		query string
// 	)

// 	if len(req.Fields) == 0 {
// 		err = errors.New("no updates provided")
// 		return
// 	}

// 	req.Fields["id"] = req.Id

// 	for key := range req.Fields {
// 		set += fmt.Sprintf(" %s = :%s ", key, key)
// 		if ind != len(req.Fields)-1 {
// 			set += ", "
// 		}
// 		ind++
// 	}

// 	query = `
// 		UPDATE
// 			"user"
// 	` + set + ` , updated_at = now()
// 		WHERE
// 			id = :id
// 	`

// 	query, args := helper.ReplaceQueryParams(query, req.Fields)

// 	result, err := u.db.Exec(ctx, query, args...)
// 	if err != nil {
// 		return
// 	}

// 	return result.RowsAffected(), err
// }

// func (u *userRepo) Delete(ctx context.Context, req *auth_service.UserPK) error {

// 	query := `DELETE FROM "user" WHERE id = $1`

// 	_, err := u.db.Exec(ctx, query, req.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

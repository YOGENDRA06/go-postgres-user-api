package repository

import (
	"context"
	"time"

	"Go_Backend_Development_Task/db/sqlc"
)

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.User, error) {
	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int64,
) (sqlc.User, error) {
	return r.q.GetUserByID(ctx, int32(id))
}

func (r *UserRepository) GetUsers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]sqlc.User, error) {

	return r.q.GetUsers(ctx, sqlc.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (sqlc.User, error) {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   int32(id),
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int64,
) error {
	return r.q.DeleteUser(ctx, int32(id))
}

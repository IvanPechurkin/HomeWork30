package interfaces

import (
	"30/pkg/model"
	"context"
)

type Repository interface {
	AddUser(ctx context.Context, user *model.User)
	DeleteUser(ctx context.Context, userId int) error
	Get(ctx context.Context, userId int) (*model.User, error)
	GetAll(ctx context.Context) []*model.User
	GetFriends(ctx context.Context, userId int) ([]*model.User, error)
	UpdateUserAge(ctx context.Context, userId int, age int) error
	LinkUsers(ctx context.Context, userLinkFrom int, userLinkTo int) error
}

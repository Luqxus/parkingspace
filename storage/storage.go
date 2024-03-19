package storage

import (
	"context"
	"github.com/luquxSentinel/spacedrive/types"
)

type Storage interface {
	CountEmail(ctx context.Context, email string) (int64, error)
	CreateUser(ctx context.Context, user *types.User) error
	GetUserWithEmail(ctx context.Context, email string) (*types.User, error)
}


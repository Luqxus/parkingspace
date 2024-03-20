package storage

import (
	"context"

	"github.com/luquxSentinel/spacedrive/types"
	"gorm.io/gorm"
)

type Storage interface {
	CountEmail(ctx context.Context, email string) (int64, error)
	CreateUser(ctx context.Context, user *types.User) error
	GetUserWithEmail(ctx context.Context, email string) (*types.User, error)
}

type storage struct {
	db gorm.DB
}

func New() *storage {
	return &storage{}
}

func (s *storage) CountEmail(ctx context.Context, email string) (int64, error) {
	var count int64

	tx := s.db.Model(&types.User{}).Where("user.email = ?", email).Count(&count)

	return count, tx.Error
}

func (s *storage) CreateUser(ctx context.Context, user *types.User) error {
	tx := s.db.Create(user)
	return tx.Error
}

func (s *storage) GetUserWithEmail(ctx context.Context, email string) (*types.User, error) {

	tx := s.db.Where("email = ?", email)

	if tx.Error != nil {
		return nil, tx.Error
	}

	user := new(types.User)

	tx.Scan(user)

	return user, tx.Error
}

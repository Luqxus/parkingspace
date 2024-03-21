package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/luquxSentinel/spacedrive/storage"
	"github.com/luquxSentinel/spacedrive/tokens"
	"github.com/luquxSentinel/spacedrive/types"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type AuthService interface {
	CreateUser(ctx context.Context, data *types.CreateUserData) error
	LoginUser(ctx context.Context, data *types.LoginData) (*types.User, string, error)
}

type authService struct {
	storage storage.AuthStorage
}

func NewAuthService(authstorage storage.AuthStorage) *authService {
	return &authService{
		storage: authstorage,
	}
}

func (s *authService) CreateUser(ctx context.Context, data *types.CreateUserData) error {
	// TODO: create a new user

	// TODO: check if no user is associated with account
	emailCount, err := s.storage.CountEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	if emailCount > 0 {
		return errors.New("email already in use")
	}
	// TODO: new user from data
	newUser := new(types.User)

	// generate new user id
	newUser.UID = uuid.NewString()
	newUser.Email = data.Email
	newUser.FirstName = data.FirstName
	newUser.LastName = data.LastName
	newUser.LastSignIn = time.Now().Local()
	newUser.IsEmailValified = false

	// hash password
	newUser.Password, err = HashPassword(data.Password)
	if err != nil {
		return err
	}

	// set created at time
	newUser.CreatedAt = time.Now().Local()

	// persist user in database
	return s.storage.CreateUser(ctx, newUser)
}

// get user by email and matching password
func (s *authService) LoginUser(ctx context.Context, data *types.LoginData) (*types.User, string, error) {
	// sign user in

	// TODO: fetch user by email
	user, err := s.storage.GetUserWithEmail(ctx, data.Email)
	if err != nil {
		log.Printf("failed to fetch user with email. error : %v", err)
		return nil, "", errors.New("wrong email or password")
	}

	// TODO: verify user password
	if err := verifyPassword(user.Password, data.Password); err != nil {
		return nil, "", errors.New("wrong email or password")
	}

	// TODO: generate jwt
	signedToken, err := tokens.GenerateJWT(user.UID, user.Email)
	if err != nil {
		log.Panic(err)
		return nil, "", err
	}

	// update last signin to time.Now().Local on successful signin
	s.updateLastSignIn(ctx, user.Email)

	// return user, jwt, error
	return user, signedToken, nil
}

func (s *authService) DeleteUser() {
	// TODO: delete user account logic
}

func (s *authService) UpdateUser() {
	// TODO: update user logic
}

func (s *authService) updateLastSignIn(ctx context.Context, email string) error {
	// update last signin to time.Now().Local on successful signin
	return s.storage.UpdateLastSignIn(ctx, email)
}

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(b), err
}

func verifyPassword(foundPassword, givenPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(foundPassword), []byte(givenPassword))
}

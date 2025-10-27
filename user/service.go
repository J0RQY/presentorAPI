package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *UserFull) error
}

type UserService struct {
	store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) CreateUser(email string, password string) (string, error) {

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	hashedPasswordStr := string(hashedPassword)

	userUUID := uuid.New()

	user := &UserFull{
		UUID:            userUUID,
		Email:           email,
		EmailVerified:   false,
		PasswordHash:    &hashedPasswordStr,
		AuthMethod:      "PASSWORD",
		OAuthProvider:   nil,
		OAuthProviderID: nil,
	}

	if err := s.store.CreateUser(context.Background(), user); err != nil {
		return "", fmt.Errorf("failed to create user in database: %w", err)
	}

	fmt.Printf("Created user with UUID: %s\n", userUUID)
	return userUUID.String(), nil
}

func hashPassword(plainPassword string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error generating bcrypt hash: %w", err)
	}
	return hashedPassword, nil
}

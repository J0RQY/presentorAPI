package userstore

import (
	"context"
	"fmt"

	"github.com/j0rqy/presentorAPI/database"
	"github.com/j0rqy/presentorAPI/user"
)

type PostgresStore struct{}

func NewPostgresStore() *PostgresStore {
	return &PostgresStore{}
}

func (s *PostgresStore) CreateUser(ctx context.Context, user *user.UserFull) error {
	db := database.GetDB()

	query := `
		INSERT INTO users (
			uuid, email, email_verified, password_hash, auth_method,
			oauth_provider, oauth_provider_id, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
	`

	_, err := db.Exec(
		ctx, query,
		user.UUID,
		user.Email,
		user.EmailVerified,
		user.PasswordHash,
		user.AuthMethod,
		user.OAuthProvider,
		user.OAuthProviderID,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

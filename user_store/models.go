package user_store

import (
	"github.com/google/uuid"
)

type UserFull struct {
	UUID            uuid.UUID
	Email           string
	EmailVerified   bool
	PasswordHash    *string
	AuthMethod      string
	OAuthProvider   *string
	OAuthProviderID *string
}

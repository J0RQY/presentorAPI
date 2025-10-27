package user

import (
	"github.com/google/uuid"
)

type CreationSuccessResponse struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

type UserCreateRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserFull struct {
	UUID            uuid.UUID
	Email           string
	EmailVerified   bool
	PasswordHash    *string
	AuthMethod      string
	OAuthProvider   *string
	OAuthProviderID *string
}

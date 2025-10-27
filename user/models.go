package user

type CreationSuccessResponse struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

type UserCreateRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

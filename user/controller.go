package user

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/j0rqy/presentorAPI/web/validation"
)

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{service: service}
}

func (h *UserController) CreateUserHandler(c *gin.Context) {
	var user UserCreateRequestDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, validation.ErrorResponse{
			Error:   "Validation Failed",
			Details: validation.FormatValidationErrors(err),
		})
		return
	}

	userUUID, err := h.service.CreateUser(strings.ToLower(user.Email), user.Password)
	if err != nil {
		log.Printf("Internal error creating user: %v", err)

		c.JSON(http.StatusInternalServerError, validation.ErrorResponse{
			Error:   "Failed",
			Details: "Internal server error.",
		})
		return
	}

	c.JSON(http.StatusCreated, CreationSuccessResponse{
		Message: "success",
		UUID:    userUUID,
	})
}

package handler

import (
	"net/http"

	"github.com/felipecveiga/crud_simples_padrao/service"
	"github.com/labstack/echo/v4"
)


type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (uh *UserHandler) GetUsers(c echo.Context) error {
	users, err := uh.UserService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {"error": "Failed to fetch users"})

	}
	return c.JSON(http.StatusOK, users)
}

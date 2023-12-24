package handler

import (
	"net/http"

	"github.com/felipecveiga/crud_simples_padrao/model"
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

//BUSCA TODOS OS USUÁRIOS
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {"error": "Failed to fetch users"})

	}
	return c.JSON(http.StatusOK, users)
}

//BUSCA O USUÁRIO ESPECIFICADO PELO ID
func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("id")

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string {"error": "Usuário não encontrado"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UserUpdate(c echo.Context) error {
	userID := c.Param("id")

	var userData model.User
	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string {"error": "erro ao ler os dados da requisição"})
	}

	//CHAMAR O MÉTODO DO SERVICE PARA ATUALIZAR O USUÁRIO
	userUpdate, err := h.UserService.UserUpdate(userID, userData) 
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string {"error": "Usuário não encontrado"})
	}
	return c.JSON(http.StatusOK, userUpdate)
}
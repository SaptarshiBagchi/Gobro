package user

import (
	"net/http"

	"gobro.starter/internal/ports/services"
)

type UserControllerInterface interface {
	Greet(w http.ResponseWriter, r *http.Request)
}

type UserControllerHandler struct {
	//Will be implementing the ports
	UserService services.UserServiceInterface
}

func NewUserController(service *services.UserService) UserControllerHandler {
	return UserControllerHandler{
		UserService: service,
	}
}

// Greet implements UserControllerInterface.
func (u UserControllerHandler) Greet(w http.ResponseWriter, r *http.Request) {
	u.UserService.Greet()
}

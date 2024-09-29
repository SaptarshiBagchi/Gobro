package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"gobro.starter/internal/ports"
	"gobro.starter/internal/ports/services"
)

func SetupUserRoutes(userRouter *mux.Router, publisher *ports.MessagePublisher) {

	userService := services.GetInstance(publisher)
	var userHandler UserControllerInterface = NewUserController(userService)
	userRouter.HandleFunc("/users/hi", userHandler.Greet).Methods(http.MethodGet)
}

package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"gobro.starter/internal/ports"
	"gobro.starter/internal/ports/services"
)

func SetupUserRoutes(userRouter *mux.Router, publisher *ports.MessagePublisher) {

	// Setting up the handler
	userService := services.GetInstance(publisher)
	var userHandler UserControllerInterface = NewUserController(userService)

	// Setting up the routes
	userRouter.HandleFunc("/users/hi", userHandler.Greet).Methods(http.MethodGet)
}

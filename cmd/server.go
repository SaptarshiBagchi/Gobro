package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gobro.starter/internal/adapters/http/user"
	"gobro.starter/internal/adapters/messaging"
	"gobro.starter/internal/ports"
)

func SetupApp() {

	router := mux.NewRouter()

	//setup message broker adapter

	var kafkaPublisher ports.PublisherConfig = messaging.NewKafkaPublisher("localhost")
	//Setup messaging broker port
	publisher := ports.GetMessagingInstance(kafkaPublisher)
	//setup user routes
	user.SetupUserRoutes(router, publisher)

	//You can add your own domain specific routers by passing on the router
	fmt.Println("Starting the server")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Errorf(err.Error())
	}
}

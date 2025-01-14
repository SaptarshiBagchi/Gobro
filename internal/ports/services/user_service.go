package services

import (
	"fmt"
	"sync"

	"gobro.starter/internal/ports"
)

type UserServiceInterface interface {
	Greet()
}

type UserService struct {
	//we can manage dependencies here dependencies
	publisher *ports.MessagePublisher
}

// !-- Ensuring the singleton
var (
	userServiceInstance *UserService
	userServiceSync     sync.Once
)

func GetInstance(publisher *ports.MessagePublisher) *UserService {
	userServiceSync.Do(func() {
		userServiceInstance = &UserService{publisher: publisher}
	})
	return userServiceInstance
}

// Service functions
func (u *UserService) Greet() {
	fmt.Println("Yaay you have reached UserService")
	u.publisher.Publish("Publishing Content from UserService")
}

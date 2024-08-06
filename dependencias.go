package main

import (
	"go-demo/delivery/controller"
	"go-demo/repositories"
	"go-demo/services"
	"go-demo/utils"

	"github.com/go-redis/redis/v8"
)

// add & export dependencies between components HERE
func GetDependencies(db *redis.Client) controller.UserSessionController {
	md := utils.GetBoolEnv("MULTIDEVICE_ENABLED", false)

	userRepository := repositories.NewUsersRepository(db)
	userService := services.NewUserService(userRepository, md)
	userController := controller.NewUserSessionController(userService)

	return userController

}

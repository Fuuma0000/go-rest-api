package main

import (
	"github.com/Fuuma0000/go-rest-api/controller"
	"github.com/Fuuma0000/go-rest-api/db"
	"github.com/Fuuma0000/go-rest-api/repository"
	"github.com/Fuuma0000/go-rest-api/router"
	"github.com/Fuuma0000/go-rest-api/usecase"
	"github.com/Fuuma0000/go-rest-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}

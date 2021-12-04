package main

import (
	"golang_architecture_ddd/infra"
	"golang_architecture_ddd/infra/database"
	"golang_architecture_ddd/interface/handler"
	"golang_architecture_ddd/usecase"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	taskRepository := infra.NewTaskRepository(database.ConnectDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := echo.New()
	handler.InitRouting(e, taskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

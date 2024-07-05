package main

import (
	_ "TodoListChallenge/docs" // Para a documentação Swagger
	"TodoListChallenge/routes"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

// @title To-Do List API
// @version 1.0
// @description Esta é uma API simples de lista de tarefas.
// @host localhost:8080
// @BasePath /

func main() {
	router := routes.SetupRouter()

	// Rota para documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":8080"))
}

package routes

import (
	"TodoListChallenge/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	taskGroup := router.Group("/tasks")
	{
		taskGroup.GET("/pending", controllers.GetPendingTasks)
		taskGroup.GET("/completed", controllers.GetCompletedTasks)
		taskGroup.POST("/", controllers.AddTask)
		taskGroup.PUT("/", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
	return router
}

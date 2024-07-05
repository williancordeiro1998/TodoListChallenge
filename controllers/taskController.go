package controllers

import (
	"TodoListChallenge/models"
	"TodoListChallenge/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var taskService = services.NewTaskService()

// @Summary Get all pending tasks
// @Description Get all tasks that are pending
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task
// @Router /tasks/pending [get]
func GetPendingTasks(c *gin.Context) {
	tasks := taskService.GetPendingTasks()
	c.JSON(http.StatusOK, tasks)
}

// @Summary Get all completed tasks
// @Description Get all tasks that are completed
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task
// @Router /tasks/completed [get]
func GetCompletedTasks(c *gin.Context) {
	tasks := taskService.GetCompletedTasks()
	c.JSON(http.StatusOK, tasks)
}

// @Summary Add a new task
// @Description Add a new task to the list
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Task to add"
// @Success 200 {object} models.Task
// @Router /tasks [post]
func AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskService.AddTask(task)
	c.JSON(http.StatusOK, task)
}

// @Summary Update an existing task
// @Description Update the status of an existing task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Task to update"
// @Success 200 {object} models.Task
// @Router /tasks [put]
func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := taskService.UpdateTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {string} string "Task deleted"
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

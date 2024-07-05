package services

import (
	"TodoListChallenge/models"
	"errors"
	"strconv"
)

type TaskService struct {
	tasks  []models.Task
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make([]models.Task, 0),
		nextID: 1,
	}
}

func (s *TaskService) GetPendingTasks() []models.Task {
	var pendingTasks []models.Task
	for _, task := range s.tasks {
		if !task.Completed {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

func (s *TaskService) GetCompletedTasks() []models.Task {
	var completedTasks []models.Task
	for _, task := range s.tasks {
		if task.Completed {
			completedTasks = append(completedTasks, task)
		}
	}
	return completedTasks
}

func (s *TaskService) AddTask(task models.Task) {
	task.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, task)
}

func (s *TaskService) UpdateTask(updatedTask models.Task) error {
	for i, task := range s.tasks {
		if task.ID == updatedTask.ID {
			if task.Completed && !updatedTask.Completed {
				if task.RevertCount >= 3 {
					return errors.New("task cannot be reverted more than 3 times")
				}
				updatedTask.RevertCount = task.RevertCount + 1
			}
			s.tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) DeleteTask(id string) error {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	for i, task := range s.tasks {
		if task.ID == taskID {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

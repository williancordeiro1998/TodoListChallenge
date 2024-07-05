package models

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Responsible string `json:"responsible"`
	Email       string `json:"email"`
	Completed   bool   `json:"completed"`
	RevertCount int    `json:"revert_count"`
	Status      string
}

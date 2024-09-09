package models

import (
	"github.com/mayrusoni10/golang-ai-task-manager/internal/config"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

var db *gorm.DB = config.DB

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetTasksByUserID(userID string) ([]Task, error) {
	var tasks []Task
	if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func CreateTask(task *Task) error {
	if err := db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTask(task *Task) error {
	if err := db.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(taskID, userID string) error {
	if err := db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&Task{}).Error; err != nil {
		return err
	}
	return nil
}

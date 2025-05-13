package services

import (
	"errors"
	"time"

	"github.com/djchanahcjd/go-todo/database"
	"github.com/djchanahcjd/go-todo/models"
)

// TodoService 定义Todo服务接口
type TodoService interface {
	CreateTodo(todo *models.Todo) error
	GetTodoByID(id uint) (*models.Todo, error)
	GetAllTodos() ([]models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id uint) error
	MarkTodoAsCompleted(id uint) error
}

// DefaultTodoService 默认的Todo服务实现
type DefaultTodoService struct{}

// NewTodoService 创建TodoService实例
func NewTodoService() TodoService {
	return &DefaultTodoService{}
}

// CreateTodo 创建新的Todo任务
func (s *DefaultTodoService) CreateTodo(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	return database.DB.Create(todo).Error
}

// GetTodoByID 根据ID获取Todo任务
func (s *DefaultTodoService) GetTodoByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	result := database.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// GetAllTodos 获取所有Todo任务
func (s *DefaultTodoService) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := database.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

// UpdateTodo 更新Todo任务
func (s *DefaultTodoService) UpdateTodo(todo *models.Todo) error {
	if todo.ID == 0 {
		return errors.New("todo id is required")
	}
	
	// 只更新指定字段
	return database.DB.Model(todo).Updates(map[string]interface{}{
		"title":       todo.Title,
		"description": todo.Description,
		"due_date":    todo.DueDate,
		"updated_at":  time.Now(),
	}).Error
}

// DeleteTodo 删除Todo任务
func (s *DefaultTodoService) DeleteTodo(id uint) error {
	return database.DB.Delete(&models.Todo{}, id).Error
}

// MarkTodoAsCompleted 标记Todo任务为已完成
func (s *DefaultTodoService) MarkTodoAsCompleted(id uint) error {
	todo, err := s.GetTodoByID(id)
	if err != nil {
		return err
	}
	todo.Completed = true
	todo.UpdatedAt = time.Now()
	return database.DB.Save(todo).Error
}

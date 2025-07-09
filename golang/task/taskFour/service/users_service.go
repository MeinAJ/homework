package service

import (
	"errors"
	"fmt"
	"golang-homework/task/taskFour/middleware"
	"golang-homework/task/taskFour/model"
	"golang-homework/task/taskFour/repository"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterUser(user *model.Users) error
	GetUserByID(id uint) (*model.Users, error)
	UpdateUser(user *model.Users) error
	DeleteUser(id uint) error
	ListUsers(page, pageSize int) ([]model.Users, int64, error)
	Login(username, password string) (string, error)
}

type userService struct {
	db       *gorm.DB
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository, db *gorm.DB) UserService {
	return &userService{userRepo: userRepo, db: db}
}

func (s *userService) RegisterUser(user *model.Users) error {
	// 使用事务
	return s.db.Transaction(func(tx *gorm.DB) error {
		txUserRepo := s.userRepo.WithTransaction(tx)
		// 验证必填字段
		if user.Username == "" || user.Email == "" || user.Password == "" {
			return errors.New("username, email and password are required")
		}
		// 对密码进行md5加密
		user.Password = model.HashPassword(user.Password)
		// 创建用户
		if err := txUserRepo.Create(user); err != nil {
			return fmt.Errorf("failed to create order: %w", err)
		}
		return nil
	})
}

func (s *userService) GetUserByID(id uint) (*model.Users, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(user *model.Users) error {
	// 业务验证逻辑
	if user.ID == 0 {
		return errors.New("invalid user ID")
	}

	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *userService) ListUsers(page, pageSize int) ([]model.Users, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.userRepo.List(page, pageSize)
}

func (s *userService) Login(username string, password string) (string, error) {
	// 验证用户名和密码
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", fmt.Errorf("failed to find user: %w", err)
	}
	// 验证密码
	if !model.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}
	// 生成token
	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, nil
}

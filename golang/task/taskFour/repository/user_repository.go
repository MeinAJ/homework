package repository

import (
	"golang-homework/task/taskFour/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	WithTransaction(tx *gorm.DB) UserRepository
	Create(user *model.Users) error
	FindByID(id uint) (*model.Users, error)
	Update(user *model.Users) error
	Delete(id uint) error
	List(page, pageSize int) ([]model.Users, int64, error)
	FindByUsername(username string) (*model.Users, error)
}

type usersRepository struct {
	db *gorm.DB
}

func (r *usersRepository) WithTransaction(tx *gorm.DB) UserRepository {
	return &usersRepository{db: tx}
}

func NewUsersRepository(db *gorm.DB) UserRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) Create(user *model.Users) error {
	return r.db.Create(user).Error
}

func (r *usersRepository) FindByID(id uint) (*model.Users, error) {
	var user model.Users
	err := r.db.Preload("Profile").First(&user, id).Error
	return &user, err
}

func (r *usersRepository) Update(user *model.Users) error {
	return r.db.Save(user).Error
}

func (r *usersRepository) Delete(id uint) error {
	return r.db.Delete(&model.Users{}, id).Error
}

func (r *usersRepository) List(page, pageSize int) ([]model.Users, int64, error) {
	var users []model.Users
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&model.Users{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Profile").Offset(offset).Limit(pageSize).Find(&users).Error
	return users, total, err
}

func (r *usersRepository) FindByUsername(username string) (*model.Users, error) {
	var users model.Users
	r.db.Where("username = ?", username).Offset(0).Limit(1).Find(&users)
	return &users, nil
}

package model

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
)

// Users 定义users结构体
type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);unique;not null" json:"username"`
	Password string `gorm:"type:varchar(50);not null" json:"password"`
	Email    string `gorm:"type:varchar(50);unique;not null" json:"email"`
	Token    string `json:"token"`
}

// Posts 定义Posts结构体
type Posts struct {
	gorm.Model
	Title   string `gorm:"type:varchar(50);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	UserId  uint   `gorm:"type:int;not null" json:"user_id"`
}

// Comments 定义Comments结构体
type Comments struct {
	gorm.Model
	PostId  uint   `gorm:"index;not null" json:"post_id"`
	UserId  uint   `gorm:"index;not null" json:"user_id"`
	Content string `gorm:"type:text;not null" json:"content"`
}

func HashPassword(password string) string {
	// 实现md5加密算法
	md5Hash := md5.Sum([]byte(password))
	return hex.EncodeToString(md5Hash[:])
}

func CheckPasswordHash(password string, password2 string) bool {
	userPassword := HashPassword(password)
	if userPassword != password2 {
		return false
	}
	return true
}

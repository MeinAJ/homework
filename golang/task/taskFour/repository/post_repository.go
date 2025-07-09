package repository

import (
	"golang-homework/task/taskFour/model"
	"golang-homework/task/taskFour/response"
	"gorm.io/gorm"
	"net/http"
)

type PostRepository interface {
	// WithTransaction 定义事物方法
	WithTransaction(tx *gorm.DB) PostRepository
	// CreatePost 定义新增文章的方法
	CreatePost(post *model.Posts)
	// GetPost 定义获取单篇文章的方法
	GetPost(id uint) *model.Posts
	// UpdatePost 定义更新文章的方法
	UpdatePost(post *model.Posts)
	// DeletePost 定义删除文章的方法
	DeletePost(id uint)
}

type postRepository struct {
	db *gorm.DB
}

func (p postRepository) WithTransaction(tx *gorm.DB) PostRepository {
	return &postRepository{db: tx}
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (p postRepository) CreatePost(post *model.Posts) {
	// 新增文章
	if err := p.db.Create(post).Error; err != nil {
		panic(response.Error(http.StatusInternalServerError, "Internal server error", nil))
	}
}

func (p postRepository) GetPost(id uint) *model.Posts {
	// 获取单篇文章
	var post model.Posts
	if err := p.db.First(&post, id).Error; err != nil {
		panic(response.Error(http.StatusNotFound, "Post not found", nil))
	}
	return &post
}

func (p postRepository) UpdatePost(post *model.Posts) {
	// 更新文章
	if err := p.db.Omit("created_at").Save(post).Error; err != nil {

		panic(response.Error(http.StatusInternalServerError, "Internal server error", nil))
	}
}

func (p postRepository) DeletePost(id uint) {
	// 删除文章
	if err := p.db.Delete(&model.Posts{}, id).Error; err != nil {
		panic(response.Error(http.StatusInternalServerError, "Internal server error", nil))
	}
}

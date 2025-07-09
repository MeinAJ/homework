package app

import (
	"github.com/gin-gonic/gin"
	"golang-homework/task/taskFour/controller"
	"golang-homework/task/taskFour/middleware"
	"golang-homework/task/taskFour/repository"
	"golang-homework/task/taskFour/service"
	"gorm.io/gorm"
)

type Application struct {
	DB                 *gorm.DB
	UserCtrl           *controller.UserController
	PostCtrl           *controller.PostController
	JwtAuthMiddleware  gin.HandlerFunc
	RecoveryMiddleware gin.HandlerFunc
	LoggerMiddleware   gin.HandlerFunc
}

func Initialize(db *gorm.DB) *Application {
	// 初始化所有Repository
	userRepo := repository.NewUsersRepository(db)
	postRepo := repository.NewPostRepository(db)

	// 初始化所有Service
	userService := service.NewUserService(userRepo, db)
	postService := service.NewPostService(postRepo, db)

	// 初始化所有Controller
	userCtrl := controller.NewUserController(userService)
	postCtrl := controller.NewPostController(postService)

	// 初始化Jwt
	jwtAuthMiddleware := middleware.JWTAuthMiddleware()

	// 全局异常捕获
	recoveryMiddleware := middleware.RecoveryMiddleware()

	// 初始化日志
	loggerMiddleware := middleware.LoggerMiddleware()

	return &Application{
		DB:                 db,
		UserCtrl:           userCtrl,
		PostCtrl:           postCtrl,
		JwtAuthMiddleware:  jwtAuthMiddleware,
		RecoveryMiddleware: recoveryMiddleware,
		LoggerMiddleware:   loggerMiddleware,
	}
}

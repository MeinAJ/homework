package router

import (
	"github.com/gin-gonic/gin"
	"golang-homework/task/taskFour/app"
)

func SetupRouter(app *app.Application) *gin.Engine {
	r := gin.Default()

	// API分组
	api := r.Group("/api/v1")
	{
		// 注册路由
		login := api.Group("/login")
		login.Use(app.RecoveryMiddleware)
		login.Use(app.LoggerMiddleware)
		{
			login.POST("", app.UserCtrl.Login)
		}
		// 用户路由
		user := api.Group("/users")
		user.Use(app.LoggerMiddleware)
		user.Use(app.JwtAuthMiddleware)
		user.Use(app.RecoveryMiddleware)
		{
			user.POST("/register", app.UserCtrl.Register)
		}
		// 文章路由
		post := api.Group("/posts")
		post.Use(app.LoggerMiddleware)
		post.Use(app.JwtAuthMiddleware)
		post.Use(app.RecoveryMiddleware)
		{
			post.POST("/create", app.PostCtrl.CreatePost)
			post.POST("/update", app.PostCtrl.UpdatePost)
			post.POST("/delete/:id", app.PostCtrl.DeletePost)
			post.POST("/list", app.PostCtrl.GetPostsByIDs)
		}
		// 健康检查路由
		r.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
		return r
	}
}

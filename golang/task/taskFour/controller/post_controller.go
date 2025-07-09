package controller

import (
	"github.com/gin-gonic/gin"
	"golang-homework/task/taskFour/model"
	"golang-homework/task/taskFour/response"
	"golang-homework/task/taskFour/service"
	"net/http"
	"strconv"
	"sync"
)

type PostController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{postService: postService}
}

func (controller *PostController) CreatePost(ctx *gin.Context) {
	var post model.Posts
	if err := ctx.ShouldBindJSON(&post); err != nil {

		panic(response.Error(http.StatusUnprocessableEntity, "Invalid post data", nil))
	}
	// 从ctx中获取用户信息
	post.UserId = ctx.MustGet("userID").(uint)
	controller.postService.CreatePost(&post)
	response.ReturnSuccess(ctx, &post)
}

func (controller *PostController) UpdatePost(ctx *gin.Context) {
	var post model.Posts
	if err := ctx.ShouldBindJSON(&post); err != nil {

		panic(response.Error(http.StatusUnprocessableEntity, "Invalid post data", nil))
	}
	// 获取文章详情
	postDetail := controller.postService.GetPost(post.ID)
	// 从ctx中获取用户信息
	post.UserId = ctx.MustGet("userID").(uint)
	// 判断是否是自己的文章
	if post.UserId != postDetail.UserId {
		panic(response.Error(http.StatusForbidden, "You are not the author of this post", nil))
	}
	// 更新文章
	controller.postService.UpdatePost(&post)
	response.ReturnSuccess(ctx, &post)
}

// DeletePost 删除文章
func (controller *PostController) DeletePost(ctx *gin.Context) {
	// 获取文章详情，将id转换为uint类型
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(response.Error(http.StatusUnprocessableEntity, "Invalid post id", nil))
	}
	postDetail := controller.postService.GetPost(uint(id))
	// 从ctx中获取用户信息
	userId := ctx.MustGet("userID").(uint)
	// 判断是否是自己的文章
	if userId != postDetail.UserId {
		panic(response.Error(http.StatusForbidden, "You are not the author of this post", nil))
	}
	// 删除文章
	controller.postService.DeletePost(uint(id))
	response.ReturnSuccess(ctx, nil)
}

// GetPostsByIDs 根据ids批量获取文章列表
func (controller *PostController) GetPostsByIDs(ctx *gin.Context) {
	var ids []uint
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		panic(response.Error(http.StatusUnprocessableEntity, "Invalid post ids", nil))
	}
	// 用于并发控制
	var wg sync.WaitGroup
	// 缓冲通道，防止阻塞
	postChannel := make(chan model.Posts, len(ids))
	// 并发获取每篇文章
	for _, id := range ids {
		wg.Add(1)
		go func(id uint) {
			defer wg.Done()
			post := controller.postService.GetPost(id)
			if post.ID != 0 {
				postChannel <- *post
			}
		}(id)
	}
	// 等待所有协程完成
	go func() {
		wg.Wait()
		close(postChannel)
	}()
	// 收集结果
	var posts []model.Posts
	for post := range postChannel {
		posts = append(posts, post)
	}
	response.ReturnSuccess(ctx, &posts)
}

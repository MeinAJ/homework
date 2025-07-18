package taskThree

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// SQL语句练习
// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

func GetConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		panic(err)
	}
	// ping
	dbPingErr := db.Ping()
	if dbPingErr != nil {
		panic(dbPingErr)
	}
	return db
}

func HandleBlog() {

	fmt.Println("handle blog")

	db := GetConnectDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// 查询所有id大于3的文章
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	queryContext, err := db.QueryContext(ctx, "select * from posts where id > ?", 3)
	if err != nil {
		return
	}
	var (
		id         int64
		title      string
		content    string
		user_id    int64
		created_at sql.NullString
		updated_at sql.NullString
		deleted_at sql.NullString
	)
	for queryContext.Next() {
		scanError := queryContext.Scan(&id, &title, &content, &user_id, &created_at, &updated_at, &deleted_at)
		if scanError != nil {
			panic(scanError)
		}
		fmt.Println("id:", id, "title:", title, "content:", content, "user_id:", user_id, "created_at:",
			created_at, "updated_at:", updated_at, "deleted_at:", deleted_at)
	}

	// 删除id = 5的文章
	result, err := db.Exec("delete from posts where id = ?", 5)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println("delete effect rows:", affected)

	// 更新id = 4的文章的标题
	result, err = db.Exec("update posts set title = ? where id = ?", "new title", 4)
	if err != nil {
		return
	}
	affected, err = result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println("update effect rows:", affected)

	// 插入一条id=6的文章
	result, err = db.Exec("insert into posts(id, title, content, user_id) values(?, ?, ?, ?)", 6, "new title", "new content", 1)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	fmt.Println("last insert id:", id)
}

// 题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

func HandleTransactions() {

	fmt.Println("handle transactions")

	db := GetConnectDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	ctx, err := db.Begin()
	if err != nil {
		return
	}

	// 事物代码，更新id=1的title，条件title=oldTitle，之后插入一条新的记录id=7，如果更新失败，则回滚事务
	exec, err := ctx.Exec("update posts set title = ? where id = ? and title = ?", "new title", 1, "old title")
	if err != nil {
		_ = ctx.Rollback()
		return
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		_ = ctx.Rollback()
		return
	}
	if affected == 0 {
		_ = ctx.Rollback()
		return
	}
	fmt.Println("update affected:", affected)

	// 插入一条id=7的文章
	exec, err = ctx.Exec("insert into posts(id, title, content, user_id) values(?, ?, ?, ?)", 7, "new title", "new content", 1)
	if err != nil {
		_ = ctx.Rollback()
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		_ = ctx.Rollback()
		return
	}
	fmt.Println("last insert id:", id)

	commitError := ctx.Commit()
	if commitError != nil {
		_ = ctx.Rollback()
		return
	}

}

// Sqlx入门
// 题目1：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Posts struct {
	Id        int64          `db:"id"`
	Title     string         `db:"title"`
	Context   string         `db:"content"`
	UserId    int64          `db:"user_id"`
	CreatedAt sql.NullString `db:"created_at"`
	UpdatedAt sql.NullString `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}

// GetConnectSqlxDB 获取sqlx的db
func GetConnectSqlxDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		panic(err)
	}
	// ping
	dbPingErr := db.Ping()
	if dbPingErr != nil {
		panic(dbPingErr)
	}
	return db
}

func HandleSqlx() {
	db := GetConnectSqlxDB()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	posts := &Posts{}
	err := db.Get(posts, "SELECT * FROM posts WHERE id = ?", 7)
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", posts.Id, "title:", posts.Title, "context:", posts.Context, "user_id:",
		posts.UserId, "created_at:", posts.CreatedAt, "updated_at:", posts.UpdatedAt, "deleted_at:",
		posts.DeletedAt)
}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

func HandleSqlxComplex() {
	db := GetConnectSqlxDB()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	var posts []Posts
	err := db.Select(&posts, "SELECT * FROM posts WHERE id > ?", 1)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		fmt.Println("id:", post.Id, "title:", post.Title, "context:", post.Context, "user_id:",
			post.UserId, "created_at:", post.CreatedAt, "updated_at:", post.UpdatedAt, "deleted_at:", post.DeletedAt)
	}
}

// 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

func MigrateGorm() *gorm.DB { // 连接数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	// 自动迁移
	err = db.AutoMigrate(&GormUsers{}, &GormPosts{}, &GormComments{})
	if err != nil {
		panic(err)
	}

	return db
}

type GormUsers struct {
	Id          int64          `gorm:"column:id;primary_key"`
	Username    string         `gorm:"column:username"`
	Password    string         `gorm:"column:password"`
	Email       string         `gorm:"column:email"`
	CreatedTime time.Time      `gorm:"type:datetime(3);gorm:column:created_time"` // MySQL 毫秒精度
	UpdatedTime time.Time      `gorm:"type:datetime(3);gorm:column:updated_time"` // MySQL 毫秒精度
	DeletedTime gorm.DeletedAt `gorm:"type:datetime(3);gorm:column:deleted_time"` // MySQL 毫秒精度
	PostCount   int64          `gorm:"column:post_count"`
}

type GormPosts struct { // 文章
	Id            int64          `gorm:"column:id;primary_key"`
	Title         string         `gorm:"column:title"`
	Content       string         `gorm:"column:content"`
	UserId        int64          `gorm:"column:user_id"`
	CreatedTime   time.Time      `gorm:"type:datetime(3);gorm:column:created_time"` // MySQL 毫秒精度
	UpdatedTime   time.Time      `gorm:"type:datetime(3);gorm:column:updated_time"` // MySQL 毫秒精度
	DeletedTime   gorm.DeletedAt `gorm:"type:datetime(3);gorm:column:deleted_time"` // MySQL 毫秒精度
	GormComments  []GormComments `gorm:"-"`
	CommentCount  int64          `gorm:"column:comment_count"`
	CommentStatus int64          `gorm:"column:comment_status"`
}

type GormComments struct {
	Id          int64          `gorm:"column:id;primary_key"`
	Content     string         `gorm:"column:content"`
	PostId      int64          `gorm:"column:post_id"`
	UserId      int64          `gorm:"column:user_id"`
	CreatedTime time.Time      `gorm:"type:datetime(3);gorm:column:created_time"` // MySQL 毫秒精度
	UpdatedTime time.Time      `gorm:"type:datetime(3);gorm:column:updated_time"` // MySQL 毫秒精度
	DeletedTime gorm.DeletedAt `gorm:"type:datetime(3);gorm:column:deleted_time"` // MySQL 毫秒精度
}

type IdCount struct {
	PostId int64
	Count  int64
}

// 题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

func GormQuery() {

	db := MigrateGorm()
	var gormPosts []GormPosts
	// 某个用户发布的所有文章
	db.Select("*").Where("user_id = ?", 1).Find(&gormPosts)
	var postIds []int64
	fmt.Println("Gorm multi Query")
	for _, gormPost := range gormPosts {
		fmt.Println("id:", gormPost.Id, "title:", gormPost.Title, "content:", gormPost.Content,
			"userId:", gormPost.UserId, "createdTime:", gormPost.CreatedTime.String,
			"updatedTime:", gormPost.UpdatedTime.String, "deletedTime:", gormPost.DeletedTime)
		postIds = append(postIds, gormPost.Id)
	}
	// 及其对应的评论信息
	var gormComments []GormComments
	var postIdAndCommentMap = make(map[int64][]GormComments)
	db.Select("*").Where("post_id in ?", postIds).Find(&gormComments)
	for _, gormComment := range gormComments {
		fmt.Println("id", gormComment.Id, "content:", gormComment.Content, "postId:", gormComment.PostId,
			"createdTime:", gormComment.CreatedTime, "updatedTime:", gormComment.UpdatedTime.String,
			"deletedTime:", gormComment.DeletedTime)
		postIdAndCommentMap[gormComment.PostId] = append(postIdAndCommentMap[gormComment.PostId], gormComment)
	}
	// 组合数据
	for index, gormPost := range gormPosts {
		gormPosts[index].GormComments = postIdAndCommentMap[gormPost.Id]
	}
	fmt.Println("postIdAndCommentMap:", postIdAndCommentMap)

	// 评论数量最多的文章信息
	var idCount IdCount
	db.Table("gorm_comments").Select("post_id, COUNT(*) as count").Group("post_id").
		Order("count DESC").Limit(1).Scan(&idCount)
	fmt.Println("post_id:", idCount.PostId, "count:", idCount.Count)

	// 找到了评论数量最多的文章，查询出文章信息
	var gormPost GormPosts
	db.Where("id = ?", idCount.PostId).First(&gormPost)
	fmt.Println("postId:", gormPost.Id, "title:", gormPost.Title, "content:", gormPost.Content,
		"userId:", gormPost.UserId, "createdTime:", gormPost.CreatedTime.String,
		"updatedTime:", gormPost.UpdatedTime.String, "deletedTime:", gormPost.DeletedTime)

}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func GormHook() {
	db := MigrateGorm()
	// 创建文章（自动触发用户计数更新）
	newPost := GormPosts{
		UserId:  1,
		Title:   "GORM Hooks",
		Content: "...",
	}
	db.Create(&newPost)
	// 删除评论（自动触发状态更新）
	db.Delete(&GormComments{Id: 2, PostId: 2, UserId: 2})
}

// AfterCreate 文章创建后的钩子函数
func (p *GormPosts) AfterCreate(tx *gorm.DB) (err error) {
	// 更新相应用户的文章数量
	result := tx.Model(&GormUsers{}).Where("id = ?", p.UserId).
		Update("posts_count", gorm.Expr("posts_count + ?", 1))
	return result.Error
}

// AfterDelete 评论删除后的钩子函数
func (c *GormComments) AfterDelete(tx *gorm.DB) (err error) {
	// 检查当前文章的剩余评论数量
	var count int64
	tx.Model(&GormComments{}).Where("post_id = ?", c.PostId).Count(&count)

	// 如果评论数为0，更新文章状态
	if count == 0 {
		result := tx.Model(&GormPosts{}).Where("id = ?", c.PostId).
			Update("comment_status", 0)
		return result.Error
	}
	return nil
}

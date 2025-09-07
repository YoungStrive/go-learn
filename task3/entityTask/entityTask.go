package entityTask

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Age      uint8
	Birthday time.Time
	gorm.Model
	//一个用户可以写很多文章一对多
	PostList []Post `gorm:"foreignKey:RefUserID"`
}

// Post 文章
type Post struct {
	ID uint
	//标题
	Title string
	//内容
	Content string
	//关联的用户id
	RefUserID uint
	//文章会有多个的评论
	CommentList []PostComment
	gorm.Model
}

// PostComment  评论
type PostComment struct {
	ID uint
	//内容
	Comment string
	//文章id
	PostId uint

	gorm.Model
}

func Run(db *gorm.DB) {
	//需要用指针
	db.Debug().AutoMigrate(&User{}, &Post{}, &PostComment{})
}

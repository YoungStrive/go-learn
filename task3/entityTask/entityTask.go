package entityTask

import (
	"fmt"
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
	//文章统计数
	PostCount uint8
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
	//评论状态
	CommentStatus string
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

// AfterCreate 钩子函数创建文章之后
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	postCount, ok := tx.InstanceGet("postCount")
	postCountNum := postCount.(int64)
	//成功
	if ok {
		fmt.Printf("postCount==%d", postCountNum)
		postCountNum++
		tx.Model(&User{}).Where(&User{ID: p.RefUserID}).Update("PostCount", postCountNum)
	}

	return
}

// 在同一个事务中更新数据
func (u *PostComment) AfterDelete(tx *gorm.DB) (err error) {
	postId := u.PostId

	var postCommentCount int64
	//查询这个文章下的评论数
	tx.Debug().Model(&PostComment{}).Where(&PostComment{PostId: postId}).Count(&postCommentCount)
	//如果为0  就把状态改成no
	if postCommentCount == 0 {
		tx.Model(&Post{}).Where(&Post{ID: postId}).Update("comment_status", "no")
	}
	return
}
func Run(db *gorm.DB) {
	//需要用指针 创建表
	//db.Debug().AutoMigrate(&User{}, &Post{}, &PostComment{})

	//保存数据
	//user := User{
	//	Name: "LiSi", Age: 20, Birthday: time.Now(),
	//	PostList: []Post{{Title: "标题1", Content: "如何上岸",
	//		CommentList: []PostComment{{Comment: "要上岸,先修路"},
	//			{Comment: "马无野草不肥"}}},
	//	},
	//}
	//
	//user1 := User{
	//	Name: "ZhanSan", Age: 18, Birthday: time.Now(),
	//	PostList: []Post{{Title: "我是张三", Content: "张三的修行之路",
	//		CommentList: []PostComment{{Comment: "事虽难，做则可成！"}}},
	//	},
	//}
	//users := []*User{&user, &user1}
	//保存数据
	//db.Debug().Create(users)
	var findUser User
	findUser.ID = 1
	////查询用户id 为1发布的所有文章及其对应的评论信息。
	//db.Debug().Model(&findUser).Preload("CommentList").Association("PostList").
	//	Find(&findUser.PostList)
	//var count int64
	//
	//db.Debug().Model(&PostComment{}).
	//	Group("PostId").Count(&count)
	//	var resultPost Post
	//查询评论数量最多的文章信息。
	//db.Table("posts").
	//	Select("posts.*").
	//	Joins("inner join (select post_id, count(*) as commentCount from post_comments " +
	//		"group by post_id order by commentCount desc limit 1) as sub on posts.id = sub.post_id").
	//	Scan(&resultPost)
	//fmt.Print(resultPost.Title)
	//
	//var postCount int64
	////查询用户id 为1的 文章数
	//db.Debug().Model(&Post{}).Where(&Post{RefUserID: findUser.ID}).Count(&postCount)
	//fmt.Printf("保存之前的用户的文章数，%d", postCount)
	//addPost := Post{Title: "我是新加的哦", Content: "内容暂时保密", RefUserID: findUser.ID}
	//
	////传一个参数过去
	//db.Debug().InstanceSet("postCount", postCount).Create(&addPost)
	//var postCountAfter int64
	//
	//db.Debug().Model(&User{}).Select("post_count").Where(&User{ID: findUser.ID}).
	//	Find(&postCountAfter)
	//
	//fmt.Printf("保存之后的用户的文章数，%d", postCountAfter)

	var findPostComment PostComment
	db.Debug().Model(&PostComment{}).First(&PostComment{}, 2).Find(&findPostComment)
	db.Debug().Delete(&findPostComment)

}

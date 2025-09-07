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
	//需要用指针 创建表
	//db.Debug().AutoMigrate(&User{}, &Post{}, &PostComment{})

	//保存数据
	user := User{
		Name: "LiSi", Age: 20, Birthday: time.Now(),
		PostList: []Post{{Title: "标题1", Content: "如何上岸",
			CommentList: []PostComment{{Comment: "要上岸,先修路"},
				{Comment: "马无野草不肥"}}},
		},
	}

	user1 := User{
		Name: "ZhanSan", Age: 18, Birthday: time.Now(),
		PostList: []Post{{Title: "我是张三", Content: "张三的修行之路",
			CommentList: []PostComment{{Comment: "事虽难，做则可成！"}}},
		},
	}
	users := []*User{&user, &user1}
	fmt.Println("users:", users)
	//保存数据
	//db.Debug().Create(users)
	var findUser User
	findUser.ID = 1
	//预加载
	db.Debug().Model(&findUser).Association("PostList").Find(&findUser.PostList)

}

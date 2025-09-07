package entry

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
	gorm.Model
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})

	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Post{})
	//db.AutoMigrate(&PostComment{})
	//单个增加
	//user := User{Name: "LiSi", Age: 20, Birthday: time.Now()}
	//users := []*User{
	//	{Name: "Jinzhu", Age: 17, Birthday: time.Now()},
	//	{Name: "Jackson", Age: 19, Birthday: time.Now()},
	//}
	//result := db.Create(&user)
	//冲突后更新
	//result := db.Clauses(clause.OnConflict{
	//	UpdateAll: true,
	//}).Create(&user)
	// 获取第一条记录（主键升序）
	//result := db.Debug().First(&User{})
	//获取一条记录，没有指定排序字段
	//db.Debug().Take(&User{})
	////查找最近的一个  根据主键降序
	//db.Debug().Last(&User{})
	////查询id为1的
	//db.Debug().First(&User{}, 1)
	//查所有
	//users := []*User{}
	//db.Debug().Find(&users)
	//db.Debug().Where("name=?", "zhangSan").First(&User{})
	user := User{}
	user.ID = 2
	db.Debug().Model(&user).Update("name", "zhujun2")

}

package main

import (
	"fmt"
	"github.com/user/pkg1/task3/entityTask"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:YL123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		fmt.Print(db)
	}
	entityTask.Run(db)

}

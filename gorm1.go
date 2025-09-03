// 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
// Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户
type User struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Posts    []Post
}

// 文章
type Post struct {
	ID       int    `gorm:"primary_key"`
	Title    string `gorm:"size:255"`
	Content  string `gorm:"size:255"`
	UserID   int
	User     User
	Comments []Comment
}

// 评论
type Comment struct {
	ID      int    `gorm:"primary_key"`
	Content string `gorm:"size:255"`
	PostID  int    `gorm:"index"`
	Post    Post
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}

	//创建用户表
	db.AutoMigrate(&User{})

	//创建文章表
	db.AutoMigrate(&Post{})

	//创建评论表
	db.AutoMigrate(&Comment{})

	fmt.Println("创建表成功")

}

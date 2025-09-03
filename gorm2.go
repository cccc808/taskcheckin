// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db,err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)"))
	if err!=nil {
		fmt.Println("连接数据库失败:",err)
		return
	}

	// 查询某个用户发布的所有文章及其对应的评论信息
	db.Where("ID .= ？"，1).PreLoad("Post").PreLoad("Comments").Find(&user)

	//预加载全部
	//db.Where("ID .= ？"，1).Preload(clause.Associations).Find(&users)

	fmt.Println("ID为1的用户发布的所有文章及其对应的评论信息："user)

	// 查询c
	var post Post
	db.Model(&Post{}).Order("CommentCount DESC").First(&post)

	fmt.Println("评论数量最多的文章信息：",post)


}


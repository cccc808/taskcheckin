// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	PostCount int    //文章数量
	Post      []Post //关联的文章
}

type Post struct {
	ID           int `gorm:"primary_key"`
	UserID       int
	Status       string    //文章状态
	CommentCount int       //评论数量
	Comment      []Comment //关联的评论
}

type Comment struct {
	ID      int `gorm:"primary_key"`
	PostID  int
	Post    Post   //关联的文章
	Content string //评论内容
}

// 为Post模型添加创建前的钩子函数
// 在文章创建时自动更新用户的文章数量统计
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 文章创建时自动更新用户的文章数量统计字段
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// 为Comment模型添加创建前的钩子函数
// 在评论创建时自动更新文章的评论数量统计
func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	err1 := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err1 != nil {
		return err1
	}

	// 评论创建时自动更新文章的评论状态为 "有评论"
	err2 := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("status", "有评论").Error
	if err2 != nil {
		return err2
	}

	return nil
}

// 为Comment模型添加删除前的钩子函数
// 在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	//先获取关联的文章
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return err
	}

	//检查评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}

	post.CommentCount = int(commentCount)
	if post.CommentCount == 0 {
		post.Status = "无评论"
	} else {
		post.Status = "有评论"
	}

	return tx.Save(&post).Error
}
func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("无法连接数据库:%v", err)
		return
	}

	//创建用户
	var user User
	result := db.Create(&User{})
	if result.Error != nil {
		log.Printf("创建用户失败:%v", result.Error)
	} else {
		db.First(&user) //获取刚创建的用户
		log.Printf("创建用户成功:ID=%d,初始文章数=%d", user.ID, user.PostCount)
	}

	//创建文章(会触发Post的BeforeCreate钩子函数)
	post := Post{UserID: user.ID}
	result = db.Create(&post)
	if result.Error != nil {
		log.Printf("创建文章失败:%v", result.Error)
	} else {
		var updatedUser User
		db.First(&updatedUser, user.ID) //获取更新后的用户
		log.Printf("创建文章成功:ID=%d,用户ID=%d,用户文章数=%d,评论数=%d", post.ID, post.UserID, updatedUser.PostCount, post.CommentCount)
	}

	//创建评论(会触发Comment的AfterDelete钩子函数)
	comment := Comment{PostID: post.ID, Content: "评论内容1"}
	result = db.Create(&comment)
	if result.Error != nil {
		log.Printf("创建评论失败:%v", result.Error)
	} else {
		var updatedPost Post
		db.First(&updatedPost, post.ID) //获取更新后的文章
		log.Printf("创建评论成功:ID=%d,文章ID=%d,评论数=%d,文章状态=%s", comment.ID, comment.PostID, updatedPost.CommentCount, updatedPost.Status)
	}

	result = db.Delete(&comment)
	if result.Error != nil {
		log.Printf("删除评论失败:%v", result.Error)
	} else {
		var updatedPost Post
		db.First(&updatedPost, post.ID) //获取更新后的文章
		log.Printf("删除评论成功:文章ID=%d,评论数=%d,文章状态=%s", comment.PostID, updatedPost.CommentCount, updatedPost.Status)
	}
}

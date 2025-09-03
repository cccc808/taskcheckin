// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、
// age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Age   int
	Grade string
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	//defer db.close()

	stu := Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Create(&stu)

	stus := []Student{{Name: "李四", Age: 19, Grade: "二年级"},
		{Name: "wshjad", Age: 20, Grade: "二年级"},
		{Name: "wshjah", Age: 14, Grade: "一年级"}}
	db.Create(&stus)

	stu1 := db.Where("age > ?", 18).Find(&stus)

	fmt.Println(stu1)

	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	db.Where("age < ?", 15).Delete(&Student{})
}

// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、
// department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
// 并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个
// Employee 结构体中。
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

var Db *sqlx.DB

// 初始化连接数据库
func Init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}

	Db = database
}

func main() {
	Init()
	//使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
	r, err := Db.Queryx("SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		fmt.Println(err)
		return
	}

	var employees []Employee
	for r.Next() {
		var employee Employee
		err = r.StructScan(&employee)
		if err != nil {
			fmt.Println(err)
			return
		}
		employees = append(employees, employee)
	}
	fmt.Println("技术部员工信息:", employees)

	//使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
	var employee Employee
	err = Db.Get(&employee, "SELECT * FROM employees WHERE salary = (SELECT MAX(salary) FROM employees)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("最高工资员工信息:", employee)

}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
package main

import "fmt"

//shape 接口
type Shape interface{
	Area() float64
	Perimeter() float64
}

//Rectangle 结构体
type Rectangle struct {
	width, height float64
}

//Circle 结构体
type Circle struct {
	radius float64
}

//实现Rectangle的Shape接口
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

//实现Rectangle的Perimeter接口
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

//实现Circle的Shape接口
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

//实现Circle的Perimeter接口
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	//创建Rectangle实例
	r := Rectangle{width: 10, height: 5}

	//创建Circle实例
	c := Circle{radius: 3}

	//调用Rectangle的Area()和Perimeter()方法
	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Rectangle Perimeter:", r.Perimeter())

	//调用Circle的Area()和Perimeter()方法
	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perimeter:", c.Perimeter())
}
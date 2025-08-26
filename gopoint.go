// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
//
//	然后在主函数中调用该函数并输出修改后的值。
package main

import "fmt"

func addten(num *int) {
	*num += 10
}

func main() {
	var mainnum int = 1
	fmt.Println("before addten:", mainnum)
	addten(&mainnum)
	fmt.Println("after addten:", mainnum)
}

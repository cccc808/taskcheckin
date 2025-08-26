// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
package main

import "fmt"

func plustow(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("Before doubling: ", nums)
	plustow(&nums)
	fmt.Println("After doubling: ", nums)
}

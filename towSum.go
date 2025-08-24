// 两数之和

// 考察：数组遍历、map使用

// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数,返回索引+值

package main

import "fmt"

func towSum(nums []int, target int) []map[int]int {
	result := make([]map[int]int, 0)
	//temp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				temp := make(map[int]int)
				temp[i] = nums[i]
				temp[j] = nums[j]
				result = append(result, temp)

			}

		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15, 3, 5, 8, 10, 6, 16}
	target := 19
	result := towSum(nums, target)
	fmt.Println("target:", target, "result:", result)
}

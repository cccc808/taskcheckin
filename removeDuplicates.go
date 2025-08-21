package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	// nums = append(nums,1)
	// nums = append(nums,3)
	// nums[1] = 9
	// fmt.Println(nums)
	// return len(nums)
	var count int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				nums[i+1] = nums[j]
				i++

			} else {
				count++
			}
		}

	}

	nums = nums[:len(nums)-count]
	fmt.Println("After removing duplicates nums in function:", nums)

	return len(nums)
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 4, 5, 6, 7, 7, 8, 9, 10}
	fmt.Println("Original array:", nums)
	length := removeDuplicates(nums)
	fmt.Println("Length of the array after removing duplicates:", length)
	fmt.Println("Original array After removing duplicates:", nums)
	fmt.Println("Needed array After removing duplicates:", nums[:length])

}

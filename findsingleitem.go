package main

import "fmt"

func findsingleitem(arr []int) []int {
	mapp := make(map[int]int)
	for _, val := range arr {
		mapp[val]++
	}
	fmt.Println("mapp=", mapp)
	var singleitem []int
	for key, count := range mapp {
		if count == 1 {
			singleitem = append(singleitem, key)
		}
	}

	return singleitem
}
func main() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 9, 9, 10, 10}

	singleitem := findsingleitem(arr)
	fmt.Println("singleitem: ", singleitem)

}

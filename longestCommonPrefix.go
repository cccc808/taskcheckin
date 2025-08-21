package main

import (
	"fmt"
)

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	for i := 0; i < minLength; i++ {
		c := strs[0][i]
		for _, str := range strs {
			if str[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLength]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Longest Common Prefix:", longestCommonPrefix(strs)) // Output: "fl"
}

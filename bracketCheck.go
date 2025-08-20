package main

import (
	"fmt"
)

func match(left, right rune) bool {
	return left == '(' && right == ')' || left == '{' && right == '}' || left == '[' && right == ']'
}

func isValid(s string) bool {
	var stack []rune

	for _, char := range s {
		switch char {
		//如果是左括号，则压入栈中
		case '(', '{', '[':
			stack = append(stack, char)
		// 如果是右括号，检查栈是否为空，或者栈顶的左括号是否与当前右括号匹配
		case ')', '}', ']':
			if len(stack) == 0 || !match(stack[len(stack)-1], char) {
				return false
			}
			//匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}

	//如果所有括号都正确匹配，栈中应该为空，返回true
	return len(stack) == 0
}
func main() {
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("()[]{"))  // false
	fmt.Println(isValid("(]"))     // false

}

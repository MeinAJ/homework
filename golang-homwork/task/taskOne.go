package task

import (
	"fmt"
	"strconv"
)

// FindNumberOfSingleNumber 只出现一次的数字 https://leetcode.cn/problems/single-number/description/
func FindNumberOfSingleNumber(nums []int) int {
	var result = make(map[int]int)
	if nums != nil && len(nums) > 0 {
		for _, num := range nums {
			if result[num] == 0 {
				// 说明第一次出现
				result[num] = 1
			} else {
				// 第二及出现
				result[num] += 1
			}
		}
		for key, value := range result {
			if value == 1 {
				fmt.Println("只出现一次的数字：", key)
				return key
			}
		}
	}
	return 0
}

// FindHuiWenShu 回文数 https://leetcode.cn/problems/palindrome-number/
func FindHuiWenShu(num int) bool {
	numStr := strconv.Itoa(num)
	for index, char1 := range numStr {
		if index <= len(numStr)/2 {
			if string(char1) != string(numStr[len(numStr)-1-index]) {
				fmt.Println("不是回文数：", num)
				return false
			}
		}
	}
	fmt.Println("是回文数：", num)
	return true
}

// IsValidBrackets 有效的括号 https://leetcode.cn/problems/valid-parentheses/description/
func IsValidBrackets(s string) bool {
	if len(s)%2 != 0 {
		fmt.Println("不是有效的括号", s)
		return false
	}
	stack := newStack(len(s))
	for _, char := range s {
		// 逻辑：每次往里面插入时，判定一下栈顶元素是否相等，相等就出栈前一个元素
		if stack.Peek() == '(' && char == ')' {
			stack.Pop()
			continue
		}
		if stack.Peek() == '{' && char == '}' {
			stack.Pop()
			continue
		}
		if stack.Peek() == '[' && char == ']' {
			stack.Pop()
			continue
		}
		stack.Push(char)
	}
	if stack.nextIndex == 0 {
		fmt.Println("是有效括号：", s)
		return true
	} else {
		fmt.Println("不是有效括号：", s)
		return false
	}
}

// Stack 实现栈的功能，先进后出FILO
type Stack struct {
	nextIndex int
	items     []interface{}
}

func newStack(size int) *Stack {
	return &Stack{nextIndex: 0, items: make([]interface{}, size)}
}

// Push 入栈
func (stack *Stack) Push(s interface{}) {
	if stack.nextIndex < len(stack.items) {
		stack.items[stack.nextIndex] = s
		stack.nextIndex++
	} else {
		fmt.Println("超出栈的大小：", s)
	}
}

// Pop 出栈
func (stack *Stack) Pop() interface{} {
	if stack.nextIndex-1 >= 0 {
		currentItem := stack.items[stack.nextIndex-1]
		stack.nextIndex--
		stack.items[stack.nextIndex] = nil
		return currentItem
	} else {
		fmt.Println("没有元素了")
	}
	return nil
}

// Peek 查看栈顶元素
func (stack *Stack) Peek() interface{} {
	if stack.nextIndex-1 >= 0 {
		return stack.items[stack.nextIndex-1]
	}
	return nil
}

// 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/description/

// 加一 https://leetcode.cn/problems/plus-one/description/

// 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

// 合并区间 https://leetcode.cn/problems/merge-intervals/description/

// 两数之和 https://leetcode.cn/problems/two-sum/description/

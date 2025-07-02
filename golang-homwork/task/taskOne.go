package task

import (
	"fmt"
	"math"
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
		// 逻辑：每次往里面插入时，判定一下栈顶元素是否相等，相等就出栈元素
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

// LongestCommonPrefix 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/description/
func LongestCommonPrefix(strs []string) string {
	// 思路：先找到最小长度，再去一一对比每一个index位置的值是否相等，每次对比时发现不相等了，后面的就不用比较了
	minSize := math.MaxInt64
	for _, str := range strs {
		if len(str) < minSize {
			minSize = len(str)
		}
	}
	result := ""
	for i := 0; i < minSize; i++ {
		currentChar := string(strs[0][i])
		for _, str := range strs {
			if currentChar != string(str[i]) {
				fmt.Println("最长前缀：", result)
				return result
			}
		}
		result += currentChar
	}
	fmt.Println("最长前缀：", result)
	return result
}

// PlusOne 加一 https://leetcode.cn/problems/plus-one/description/
func PlusOne(digits []int) []int {
	// 思路先转成字符串，再转成int，再加1，再转成[]int
	plusOne := 1
	size := len(digits)
	for index := range digits {
		// 倒叙判断
		if digits[size-index-1]+plusOne == 10 {
			plusOne = 1
			digits[size-index-1] = 0
		} else {
			plusOne = 0
			digits[size-index-1] = digits[size-index-1] + 1
			break
		}
	}
	if plusOne == 1 {
		// 需要进一位
		result := make([]int, size+1)
		result[0] = 1
		for i := range digits {
			result[i+1] = digits[i]
		}
		fmt.Println("plus one:", result)
		return result
	}
	fmt.Println("plus one:", digits)
	return digits
}

// RemoveDuplicates 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func RemoveDuplicates(nums []int) int {
	// 思路：当前元素与后面所有元素对比，知道找到不同的元素，每次循环后beginIndex=beginIndex+1，找到不同元素的作为nextIndex
	size := findNext(nums, 0, 1, 1)
	fmt.Println("唯一元素个数：", size)
	fmt.Println("当前元素：", nums)
	return size
}

func findNext(nums []int, beginIndex int, nextIndex int, size int) int {
	num := nums[beginIndex]
	for i := nextIndex; i < len(nums); i++ {
		if num != nums[i] {
			// 发现了一个不同的，就不同数的数量+1
			size++
			// 不同时，设置当前不同的元素到beginIndex位置的下一位
			if beginIndex+1 < len(nums) && (beginIndex+1) < i {
				nums[beginIndex+1] = nums[i]
			}
			return findNext(nums, beginIndex+1, i+1, size)
		}
	}
	// 走到这里可以直接return了
	return size
}

// Merge 合并区间 https://leetcode.cn/problems/merge-intervals/description/
func Merge(intervals [][]int) [][]int {
	var result [][]int
	skipMap := make(map[int]bool)
	for i := 0; i < len(intervals); i++ {
		if skipMap[i] {
			continue
		}
		for j := i + 1; j < len(intervals); j++ {
			small, big := intervals[i][0], intervals[i][1]
			smallOfNext, bigOfNext := intervals[j][0], intervals[j][1]
			// 两两值比较,有三种情况。
			if big >= smallOfNext && bigOfNext >= small {
				// 存在交集，保留第一个元素，skip第二个元素,合并前后两个元素
				intervals[i][0] = int(math.Min(float64(small), float64(smallOfNext)))
				intervals[i][1] = int(math.Max(float64(big), float64(bigOfNext)))
				skipMap[j] = true
			}
		}
		result = append(result, intervals[i])
	}
	fmt.Println("合并后的区间：", result)
	return result
}

// 两数之和 https://leetcode.cn/problems/two-sum/description/

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
				fmt.Printf("只出现一次的数字：%d\n", key)
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
				fmt.Printf("不是回文数：%d\n", num)
				return false
			}
		}
	}
	fmt.Printf("是回文数：%d", num)
	return true
}

// 有效的括号 https://leetcode.cn/problems/valid-parentheses/description/
func isValid(s string) bool {
	a := "("
	b := ")"
	c := "{"
	d := "}"
	e := "["
	f := "]"

	var arrs = make([]string, len(s))
	for index, char := range s {

	}
}

// 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/description/

// 加一 https://leetcode.cn/problems/plus-one/description/

// 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

// 合并区间 https://leetcode.cn/problems/merge-intervals/description/

// 两数之和 https://leetcode.cn/problems/two-sum/description/

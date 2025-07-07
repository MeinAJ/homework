package main

import (
	"fmt"
	"golang-homework/task"
	"sync"
)

// main方法是程序的入口
// package必须声明为main，同时也必须有个main方法
// 运行程序：go run main.go
func main() {
	// 方法名大写开头，表示该方法为公开方法，可以被其他包调用
	// 方法名小写开头，表示该方法为私有方法，只能在当前包内访问

	// task1
	task.FindNumberOfSingleNumber([]int{1, 1, 2, 2, 3, 4, 4, 5, 5})
	task.FindHuiWenShu(212)
	task.IsValidBrackets("[][][{}]()(){}")
	task.LongestCommonPrefix([]string{"1234567", "123456", "123456"})
	task.PlusOne([]int{9, 9, 9, 9, 8})
	task.RemoveDuplicates([]int{0, 1, 2, 3, 3, 3, 3, 5, 6, 6, 7})
	task.Merge([][]int{{1, 3}, {2, 3}, {3, 5}, {6, 7}})
	task.TwoSum([]int{1, 2, 3, 4, 5, 6}, 30)

	// task2
	num := 1
	fmt.Println("Add10：", num)
	task.Add10(&num)
	fmt.Println("Add10：", num)

	nums := []int{1, 2, 3}
	fmt.Println("Multi2WithEveryItemOfSlice：", nums)
	task.Multi2WithEveryItemOfSlice(nums)
	fmt.Println("Multi2WithEveryItemOfSlice：", nums)

	group := sync.WaitGroup{}
	group.Add(2)
	task.PrintOddAndEvenFrom0To10(&group)
	group.Wait()

	group = sync.WaitGroup{}
	group.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			task.PrintAsync(i)
		}()
	}
	group.Wait()

	task.AtomicCounter()

	task.HandleBlog()

	task.HandleTransactions()

}

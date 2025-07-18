package taskTwo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 创建独立的随机数生成器
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Add10 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func Add10(num *int) {
	*num += 10
}

// Multi2WithEveryItemOfSlice 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func Multi2WithEveryItemOfSlice(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}
}

// PrintOddAndEvenFrom0To10 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func PrintOddAndEvenFrom0To10(group *sync.WaitGroup) {
	go func() {
		defer group.Done()
		for i := 1; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println("偶数：", i)
			}
		}
	}()

	go func() {
		defer group.Done()
		for i := 1; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println("奇数：", i)
			}
		}
	}()
}

// PrintAsync 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func PrintAsync(index int) {
	now := time.Now()
	// 创建新的随机数生成器
	sleepTime := time.Duration(rng.Intn(5)+5) * time.Millisecond
	time.Sleep(sleepTime)
	fmt.Println("执行任务：", index, "时间耗时", time.Since(now))

}

// Shape 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area(area float64) float64
	Perimeter(perimeter float64) float64
}

type Rectangle struct {
}

type Circle struct {
}

func (r Rectangle) Area(area float64) float64 {
	fmt.Println("Area is:", area)
	return area
}

func (r Rectangle) Perimeter(perimeter float64) float64 {
	fmt.Println("Perimeter is:", perimeter)
	return perimeter
}

func (c Circle) Area(area float64) float64 {
	fmt.Println("Area is:", area)
	return area
}

func (c Circle) Perimeter(perimeter float64) float64 {
	fmt.Println("Perimeter is:", perimeter)
	return perimeter
}

// Person 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("EmployeeID：", e.EmployeeID, "，Name：", e.Person.Name, "，Age", e.Person.Age)
}

func NewEmployee(Name string, Age int, EmployeeID string) *Employee {
	return &Employee{
		Person: Person{
			Name: Name,
			Age:  Age,
		},
		EmployeeID: EmployeeID,
	}
}

// GoRoutineReact 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func GoRoutineReact() {
	printChannel := make(chan int, 10)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	// 先写入
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			printChannel <- i
		}
		// 关闭通道
		close(printChannel)
		defer waitGroup.Done()

	}()
	// 再读出
	go func() {
		defer waitGroup.Done()
		for printInt := range printChannel {
			fmt.Println("printInt:", printInt, ",time:", time.Now())
		}
	}()
	waitGroup.Wait()
}

// GoRoutineReactWithTimeout 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func GoRoutineReactWithTimeout() {
	printChannel := make(chan int, 10)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	// 先写入
	go func() {
		defer waitGroup.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(200 * time.Millisecond)
			printChannel <- i
		}
		// 关闭通道
		close(printChannel)
	}()
	// 再读出
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case dd, ok := <-printChannel:
				if !ok {
					fmt.Println("Channel closed")
					return
				}
				fmt.Println("digit number:", dd, ",time:", time.Now())
			case <-ctx.Done():
				fmt.Println("time out")
				return
			}
		}
	}()
	waitGroup.Wait()
}

// SyncMutexCount 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func SyncMutexCount() {
	mutex := sync.Mutex{}
	count := 0
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go func(inc *int) {
			mutex.Lock()
			defer mutex.Unlock()
			defer waitGroup.Done()
			for i := 0; i < 1000; i++ {
				*inc++
			}
		}(&count)
	}
	waitGroup.Wait()
	fmt.Print("count:", count)
}

// AtomicCounter 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func AtomicCounter() {
	atomicCount := atomic.Int32{}
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go func(atomicCount *atomic.Int32) {
			defer waitGroup.Done()
			for i := 0; i < 1000; i++ {
				time.Sleep(1 * time.Millisecond)
				atomicCount.Add(1)
			}
		}(&atomicCount)
	}
	waitGroup.Wait()
	fmt.Println("atomicCount：", atomicCount.Load())
}

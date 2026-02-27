package main

import "fmt"

// 函数内的普通局部变量，通常分配在栈上
// 当函数执行完毕，栈空间会被立即回收
func stackVariable() {
	a := 10
	fmt.Printf("栈变量 a 的地址: %p, 值: %d\n", &a, a)
}

// 返回局部变量的指针，这会导致变量“逃逸”到堆上
// 因为函数结束后外部还需要访问它，所以不能随栈销毁
func heapVariable() *int {
	b := 20
	fmt.Printf("堆变量 b 的地址: %p, 值: %d\n", &b, b)
	return &b
}

func main() {
	fmt.Println("=== 栈分配示例 ===")
	stackVariable()
	// 再次调用，可能会复用刚才的栈空间（地址可能相同，也可能不同，取决于运行时调度）
	stackVariable()

	fmt.Println("\n=== 堆分配示例 (逃逸分析) ===")
	ptr := heapVariable()
	fmt.Printf("main中接收到的指针: %p, 值: %d\n", ptr, *ptr)
}

package main

// 这个结构体比较小，通常适合在栈上分配
type smallStruct struct {
	a, b int
}

// 1. 纯栈分配
// x 从未离开这个函数，也没有传给 interface{} (如 Println)
// 所以它理论上绝对应该在栈上
func stackAllocation() int {
	x := smallStruct{a: 10, b: 20}
	return x.a + x.b
}

// 2. 堆分配 (逃逸)
// y 的地址被返回了，虽然 y 是局部变量，
// 但为了保证函数结束后还能访问，编译器必须把它放到堆上
func heapAllocation() *smallStruct {
	y := smallStruct{a: 30, b: 40}
	return &y
}

// 3. 大对象分配 (可能逃逸到堆)
// 即使不返回，如果对象太大，栈放不下，也可能放到堆上
// 比如申请一个很大的数组
func largeAllocation() int {
	// 申请 10MB 的数组
	var largeArray [1024 * 1024 * 1024]byte
	largeArray[0] = 1
	return int(largeArray[0])
}

func main() {
	stackAllocation()
	heapAllocation()
	largeAllocation()
}

package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== 1. 数组 (Array) - 绝对连续 ===")
	arr := [3]int{10, 20, 30}
	fmt.Printf("arr[0]: %p\n", &arr[0])
	fmt.Printf("arr[1]: %p (距离上一个 %d 字节)\n", &arr[1], uintptr(unsafe.Pointer(&arr[1]))-uintptr(unsafe.Pointer(&arr[0])))
	fmt.Printf("arr[2]: %p (距离上一个 %d 字节)\n", &arr[2], uintptr(unsafe.Pointer(&arr[2]))-uintptr(unsafe.Pointer(&arr[1])))

	fmt.Println("\n=== 2. 切片 (Slice) - 底层数组连续 ===")
	slc := []int{10, 20, 30}
	fmt.Printf("slc[0]: %p\n", &slc[0])
	fmt.Printf("slc[1]: %p (距离上一个 %d 字节)\n", &slc[1], uintptr(unsafe.Pointer(&slc[1]))-uintptr(unsafe.Pointer(&slc[0])))
	// 切片本身只是一个很小的结构体头（ptr, len, cap），但它指向的数据区是连续的。

	fmt.Println("\n=== 3. 结构体 (Struct) - 字段相对连续 ===")
	type MyStruct struct {
		A int64
		B int8 // 只需要1字节，但为了对齐，后面可能有7字节填充(padding)
		C int64
	}
	s := MyStruct{A: 1, B: 2, C: 3}
	fmt.Printf("Struct A: %p\n", &s.A)
	fmt.Printf("Struct B: %p (距离A %d 字节)\n", &s.B, uintptr(unsafe.Pointer(&s.B))-uintptr(unsafe.Pointer(&s.A)))
	fmt.Printf("Struct C: %p (距离B %d 字节 - 包含填充)\n", &s.C, uintptr(unsafe.Pointer(&s.C))-uintptr(unsafe.Pointer(&s.B)))

	fmt.Println("\n=== 4. Map (映射) - 不连续 ===")
	// Map 的底层是哈希桶数组，Key-Value 对散落在不同的桶里
	m := make(map[int]int)
	m[0] = 10
	m[1] = 20
	m[2] = 30
	// 甚至不能对 map 元素取地址 &m[0] (编译器直接报错)
	// 因为 map 可能会扩容重排，地址会变，所以 Go 禁止获取 map 值的地址。
	fmt.Println("Map 元素的地址无法获取 (编译期禁止)，因为它们在内存中是散乱且不稳定的。")

	fmt.Println("\n=== 5. 链表 (Linked List) - 绝对不连续 ===")
	l := list.New()
	e1 := l.PushBack(10)
	e2 := l.PushBack(20)
	e3 := l.PushBack(30)
	
	// e1, e2, e3 是 *list.Element 指针
	// 它们指向的 Element 结构体是在堆上单独分配的，地址毫无规律
	fmt.Printf("Node 1: %p\n", e1)
	fmt.Printf("Node 2: %p (距离上一个 %d 字节 - 随机)\n", e2, uintptr(unsafe.Pointer(e2))-uintptr(unsafe.Pointer(e1)))
	fmt.Printf("Node 3: %p (距离上一个 %d 字节 - 随机)\n", e3, uintptr(unsafe.Pointer(e3))-uintptr(unsafe.Pointer(e2)))
}

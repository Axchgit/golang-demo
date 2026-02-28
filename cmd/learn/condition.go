package main

// 条件控制

import "fmt"

func main() {
	a, b := 1, 2

	// if/else
	if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}

	// 简单表达式
	if x := 1; x > 1 {
		fmt.Println(x)
	} else {
		fmt.Println("错误的")
	}

	score := 80

	if score > 90 {
		fmt.Println("优秀")
	} else if score >= 75 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// switch 分支类型需要和表达式一致
	switch_val := "b"

	switch switch_val {
	case "a":
		fmt.Println("是a")
	case "b":
		fmt.Println("是b")
		fallthrough
	case "bc":
		fmt.Println("包含b")
		fallthrough
	case "c":
		fmt.Println(switch_val)
	}

}

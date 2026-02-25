package main

import (
	"fmt"
	"math"
)

// 加法
func Add(a, b int) int {
	return a + b
}

// 减法
func Sub(a, b int) int {
	return a - b
}

// 乘法
func Mul(a, b int) int {
	return a * b
}

// 除法（返回结果和误差）
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// 计算圆的面积
func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	fmt.Println("=== 计算器 Demo ===")
	
	fmt.Printf("10 + 5 = %d\n", Add(10, 5))
	fmt.Printf("10 - 5 = %d\n", Sub(10, 5))
	fmt.Printf("10 * 5 = %d\n", Mul(10, 5))
	
	result, err := Div(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 5 = %d\n", result)
	}
	
	// 错误处理演示
	_, err = Div(10, 0)
	if err != nil {
		fmt.Println("捕获错误:", err)
	}
	
	fmt.Printf("半径为5的圆面积 = %.2f\n", CircleArea(5))
}

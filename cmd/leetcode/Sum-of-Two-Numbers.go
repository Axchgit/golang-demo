package main

import "fmt"

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

// 你可以按任意顺序返回答案。

func main() {

	fmt.Println(twoSum([]int{1, 2, 4, 5, 6}, 7))
	fmt.Println(twoSumPro([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	for index1, value1 := range nums {
		for index2, value2 := range nums {
			if index1 == index2 {
				continue
			}
			if value1+value2 == target {
				return []int{index1, index2}
			}
		}
	}
	return []int{}
}

func twoSumPro(nums []int, target int) []int {
	// 构造hash数组
	mid_map := make(map[int]int, 10)
	for index, value := range nums {
		// 如果 target - value 存在于hash中，则返回 index，hash值
		if val, exist := mid_map[target-value]; exist {
			return []int{index, val}
		}
		mid_map[value] = index
		// 把index作为值，value作为key存入hash数组
	}
	return []int{}
}

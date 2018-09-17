//采取的策略是边寻找，边创建map; 用空间换时间
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	//采取的策略是边寻找边存储
	for i, b := range nums {
		//找到符合的条件
		if j, ok := m[target-b]; ok {
			return []int{j, i}
		}
		m[b] = i
	}
	return nil
}

//入口函数
func test_1() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

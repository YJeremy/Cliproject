/*
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
*/

package algorithmStudy

import (
//"fmt"
)

func TwoSum1(nums []int, target int) []int {
	n := len(nums)
	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//对于数组“即操即用”
//借用go自身的哈希——遇到搜索问题可以直接创建Map数据类型，再使用其哈希方法啊！
func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	//fmt.Println("Sum2", m)

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			//赋值存在而已；因为是“两两相减”，一开始减数是存入内容，后面的被减数还会回来一轮查询的！
			//前面检查的查询
			return []int{j, i}
		} else {
			m[v] = i
		}
	}
	//fmt.Println("sum22", m) // 7 11 15 2
	return nil
}

//先操完再用
func TwoSum3(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		}
	}
	return nil
}

/*
func main() {
	r1 := TwoSum1([]int{2, 7, 11, 15}, 13)
	r2 := TwoSum2([]int{2, 7, 11, 15}, 10)
	r3 := TwoSum3([]int{2, 7, 11, 15}, 13)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}

*/

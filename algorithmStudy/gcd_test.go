//功能测试

package algorithmStudy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const N = 300

func init() {
	rand.Seed(time.Now().Unix())
}

//功能测试 执行go test -v
func TestGCD(t *testing.T) {
	num := GCD(200, 100)
	if num != 100 {
		t.Error("测试失败")
	}

}

func ExampleGCD() {

	fmt.Println(GCD(200, 100))
	fmt.Println(TwoSum1([]int{1, 2, 3, 4, 5}, 9))
}

//压力测试 执行 go test -bench=.
func BenchmarkGCD(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GCD(20000, 10000)
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum1(nums, 9)
	}

}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2(nums, 9)
	}

}

func BenchmarkTwoSum3(b *testing.B) {

	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum3(nums, 9)
	}

}

func TestMergeKLists(t *testing.T) {
	l1 := &Node{1, &Node{2, &Node{4, nil}}}
	l2 := &Node{1, &Node{3, &Node{4, nil}}}
	lists := []*Node{l1, l2}
	head := MergeKLists3(lists)
	println(head)
}

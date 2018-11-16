package word

import (
	"fmt"
	"testing"
)

//单元测试
func TestReverse(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	cases := []struct{ in, want string }{{"Hello,world", "dlrow,olleH"}, {"Hello,世界", "界世,olleH"}, {"", ""}}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q,want %q", c.in, got, c.want)
		}
	}
}

//示例验证 Example开头，go test的时候会自动执行这些示例
func ExampleReverse() {
	fmt.Println(Reverse("Hello, 12世界")) //世界前面要有空格，程序会自动trim，而忽略前后的空格比较
	//Output: 界世 ,olleH
}

//基准测试，执行b.N次，有对比性；b.N的值根据实际去那个看调整的，从而保证测试的稳定性。; 执行 go test -bench .
func BenchmarkReverse(b *testing.B) {
	b.ResetTimer() //这句之前的处理不会放到执行时间里，也不会输出到报告中， 可以在这之前做些不计划作为测试报告的操作
	for i := 0; i < b.N; i++ {
		Reverse("s string")
	}
}

//基准测试可以开启 并行测试，默认会以逻辑CPU个数进行并行测试，需开启b.RunParallel(func(pb *testing.PB))方法
/*func BenchmarkReverseParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Reverse("s string")
		}
	})
}*/

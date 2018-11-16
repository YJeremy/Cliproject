package algorithmStudy

import (
	"fmt"
)

func MokeyUser() {

	fmt.Println("777", MokeyBanana(777, 5))
	fmt.Println("3121", MokeyBanana(3121, 5))
	fmt.Println("3321", MokeyBanana(3321, 5))
	fmt.Println("3111", MokeyBanana(3111, 5))

	//查找最小香蕉数；复杂度估计是O（n2）
	for i := 6; ; i++ {
		if MokeyBanana(i, 5) == 0 {
			continue
		} else {
			fmt.Println("find the min Banana number is :", i)
			break //跳出for，否则会不断计算剩余符合条件的数，我们只取第一次出现的数，即最小的数；

		}

	}
}

/*入手
顺序思考：在草稿纸上大概可以知道一些列举算数
感到像函数重复的递归
但是，由于选项，且选项不是从0开始；递归次数可用 猴子次数来表示

直观地，从头开始

既然可以递归，那么使用遍历尝试

既然需要多次条件，那么可以考虑 1 递归 2 循环遍历；

如果是多次嵌套，那么考虑 嵌套for 或者分布一个函数，后者可读性都好



*/
func MokeyBanana(banana, turn int) int {
	//b为每次香蕉剩余数
	b := banana

	for i := 1; i < turn; i++ { //从第一次开始；前取5次，但是判断的是4次，第5次取走后剩余不判断是否满足（后面其实不满足了，编码后才知，所以，编码也要联系实际的业务逻辑，来做调整；而且是从中间的数来找的）

		b = b - 1 - (b-1)/5 //
		//fmt.Println("%d/n", b) 打印调试看看过程,有打印4次
		if (b-1)%5 == 0 { //剩余香蕉判断，本次满足则下一轮的计算，5次中能进行到第4次的剩余数还符合要求即可；
			continue
		} else {
			//fmt.Printf("%d", b) 这里打印看看输出结果的是一个很长的类型数，因为余数有小数相加
			return 0
		}

	}

	return banana
	//使用for遍历 还是 return来递归自身函数

}

//累加版本
func MokeyBananaAdd(n int) {

}

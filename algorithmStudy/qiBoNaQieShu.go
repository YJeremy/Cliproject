package algorithmStudy

//package main

import (
	"fmt"
	"strconv"
)

func QiUser() {
	var n int

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println("第" + strconv.Itoa(i) + "月兔子对数为：" + strconv.Itoa(Fibonacci(i)))
	}

}

func Fibonacci(n int) int {
	f0 := 1
	f1 := 1
	f2 := 1
	if n <= 2 {
		return 1
	}

	for i := 1; i <= n-2; i++ {
		f2 = f0 + f1
		f0 = f1
		f1 = f2
	}

	return f2

}

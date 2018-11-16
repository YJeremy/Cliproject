package algorithmStudy

//求两个数最大公约数
func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

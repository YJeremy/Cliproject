package word

func Reverse(s string) string {
	r := []rune(s) // rune 是int32的别名
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*
s:= "hello你好"
fmt.Println（len(s)）//输出长度11
fmt.Println(len([]rune(s)))//输出长度为7

"你好"//输出长度6 ，rune输出长度为2
“你” //输出长度为6 ，rune输出20320

rune 可以表示unicode编码的int值

*/

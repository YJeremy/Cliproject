/*
v02版本，改名为万能计算机；添加计算各种算法的计算机
添加：两个甚至多个数里面的最大公约数查找功能 GCD
V02.1 改进：把各个函数分块成包，以后在各个函数包里操作,增加p操作，名字对决游戏
V02.2 改进：加入RSS源下载关键字新闻并保存为MD格式（TO DO）
*/
package main

import (
	"fmt"
)

func main() {

	Factory() //在相同main包里有的函数
	//printExplain()
	fmt.Println("-h显示帮助")

}

//程序的说明书
func printExplain() {
	citydate := map[string]int{"广州": 101280101, "阳江": 101281801, "香港": 101320101, "澳门": 101330101, "台湾": 0}
	fmt.Println("\n*操作说明：")
	fmt.Println("命令行后面加参数“-h”,了解更多，如：", "\n", "./ApiURL -h")
	fmt.Println("====================")
	fmt.Println("例子：查询今天广州天气输入：", "\n", "./ApiURL -o w -i 101280101 -d 今天")
	fmt.Println("例子：求22和47两数的最大公约数输入：", "\n", "./ApiURL -o GCD -n1 22 -n2 47")
	fmt.Println("城市码:", citydate)
	fmt.Println("===================")
	fmt.Println("名字对决游戏：-o p -na1 名字1 -na2 名字2")
	fmt.Println("开始默认血量20,可以通过参数 -hp1 -hp2 修改两人血量")
	fmt.Println("到排序输入的内容： -o re -i 你好")
	fmt.Println("\n")

	fmt.Println("Enjoy yourself! \n")
}

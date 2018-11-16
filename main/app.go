package main

import (
	"CliProject/algorithmStudy"
	"CliProject/namegame"
	"CliProject/weatherApi"
	"CliProject/word"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func Factory() {
	//小程序实例化
	app := cli.NewApp()
	app.Name = "claculater-Jeremy"
	app.Usage = "万能计算机程序"
	app.Version = "v2.01"

	//小程序可用的命令操作符号，当app运行时候读取app.run的输入参数作为命令
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input,i",
			Value: "101280101",
			Usage: "w命令、re命令下使用，参数input 缩写i；含义：要输入处理的值“城市码”；默认值是广州的城市码：101280101",
		},
		cli.StringFlag{
			Name:  "day,d",
			Value: "今天",
			Usage: "w命令下使用，参数day 缩写d ；含义：可输入值： 今天，昨天，预测 三种日期；默认参数：今天",
		},
		cli.StringFlag{
			Name:  "Author, r",
			Value: "Jeremy",
			Usage: "命令Author 简写r ；作者名称",
		},
		cli.StringFlag{
			Name:  "operation,o",
			Value: "nil",
			Usage: `命令operation 简写o,含义操作命令命令参数可输入：GCD 表示求两个数的最小公倍数，w查询天气，p“你的名字”游戏；re 表示输入字符串的倒序排序输出；命令默认值：nil`,
		},
		cli.IntFlag{ //数值类型，值也是数值型
			Name:  "num1,n1",
			Value: 10,
			Usage: "GCD命令下使用；参数num1 简写n1 ；含义：计算的第一位数；默认值：10",
		},
		cli.IntFlag{
			Name:  "num2,n2",
			Value: 10,
			Usage: "GCD命令下使用；参数num2 简写n2； 含义：计算的第二位数；默认值：10",
		},
		cli.StringFlag{
			Name:  "name1,na1",
			Value: "无名氏1号",
			Usage: "p命令下使用，user name1 -n1 **；游戏玩家名字1",
		},
		cli.StringFlag{
			Name:  "name2,na2",
			Value: "无名氏2号",
			Usage: "p命令下使用，user name2:-n2 **；游戏玩家名字2",
		},
		cli.Int64Flag{
			Name:  "name1HP,hp1",
			Value: 20,
			Usage: "p命令下使用，user name1 HP:-hp1 20；游戏玩家1血量值设置",
		},
		cli.Int64Flag{
			Name:  "name2HP,hp2",
			Value: 20,
			Usage: "p命令下使用，user name2 HP:-hp2 20；游戏玩家2血量值设置",
		},
	}

	//小程序功能及其函数实现
	app.Action = func(c *cli.Context) error {

		//创建变量，读入对应输入的参数值内容，因此需要对不同的参数进行数据转换；而功能名不需要
		input := c.String("input")
		day := c.String("day")
		num1 := c.Int("num1")
		num2 := c.Int("num2")
		operation := c.String("operation")
		name1 := c.String("name1")
		name2 := c.String("name2")
		name1hp := c.Int("name1HP")
		name2hp := c.Int("name2HP")

		switch operation {
		case "w":
			a := weatherApi.Weather(input, day)
			return a
		case "GCD":
			fmt.Println("输入的两个数最大公约数是", algorithmStudy.GCD(num1, num2))
		case "p":
			namegame.NameGame(name1, name2, name1hp, name2hp)
		case "re":
			fmt.Println("倒序排序结果：", word.Reverse(input))

		}
		return nil

	}

	app.Run(os.Args) //小程序启动，启动读取输入命令行的功能，等候输入
}

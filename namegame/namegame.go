//V1.0 20180802
//chanl 无缓冲相互传递
package namegame

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//创建计数器对象
var wg sync.WaitGroup

//初始化随机数的初始值，否则每次程序开始都是同样的“随机值”
func init() {
	//rand.Seed() ，里面加入一个可以变化的参数，让Seed值开始每次都不一样。
	//time.Now().UnixNano()获取的是当前时间的纳秒，因为计算机运行太快了。这里并不需要连续获取随机数种子，只要每次开机运行代码的时间不同就可以了；
	rand.Seed(time.Now().UnixNano())
}

func NameGame(name1 string, name2 string, name1hp int, name2hp int) {

	court := make(chan int)

	fmt.Printf("游戏开始%s血量值为%d\t%s血量值为%d\n", name1, name1hp, name2, name2hp)
	// 增加并发计数
	wg.Add(2)

	go player(name1, court, 1, name2hp)
	go player(name2, court, 1, name1hp)

	// 给通道传入数据，触发并发开始游戏
	court <- 1

	// 等候计数器完结
	wg.Wait()

}

// 游戏主要逻辑
func player(name string, court chan int, turn, hp int) {
	// 函数结束时，关闭计数器；闭包传值
	defer wg.Done()
	skill := []string{"会心一击", "阳光烈焰", "天马流星拳", "升龙拳", "吸血光环", "十万伏特", "射火焰", "还我漂漂拳", "十字死光", "百万吨吸收", "挠痒痒 ", "狮子偷桃", "夺命剪刀脚", "龟波气功"}
	hecai := []string{"厉害厉害～", "优秀！", "skr～", "啊哟，不错哦！", "一个字，稳！"}

	for {

		battle, ok := <-court

		if !ok {

			fmt.Printf("%s输了,555～\n", name)
			return
		}

		n := rand.Intn(10) //产生给对方伤害值
		hecainum := rand.Intn(4)

		fmt.Printf("第%d回合 ", turn)
		turn++
		if n == 4 {
			hp = -1
			fmt.Printf("你的名字|%-6s|使出秘籍“要你命3000”！直接将对手血量打成0！\n%s胜利！\n", name, name)
			close(court)
			return
		}

		hp = hp - n //血量减去伤害
		if hp <= 0 {
			hp = 0
			fmt.Printf("KO！%s将对方血量打到0！ %s胜利！\n", name, name)

			// 关闭通道，退出循环
			close(court)
			return
		}
		fmt.Printf("你的名字|%-6s|使用|%s输出%d点伤害,%-6s|\t将对手血量降到%d\n", name, skill[n], n, hecai[hecainum], hp)

		time.Sleep(1 * time.Second)
		court <- battle //阻塞通道
	}
}

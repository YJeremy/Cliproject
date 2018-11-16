/*
单向链表数据结构制作
*/
package algorithmStudy

import "fmt"

type Object interface{}

//定义节点
type Node2 struct {
	data Object
	next *Node2
}

//定义单向链表
type List struct {
	head *Node2
	tail *Node2
	size uint64
}

//初始化
func (list *List) Init() {
	(*list).size = 0   //此时链表是空的
	(*list).head = nil //没有头
	(*list).tail = nil //没有尾
}

//向尾部添加数据
func (list *List) Append(node *Node2) bool {
	if node == nil {
		return false
	}

	(*node).next = nil
	//将新元素放入单链表中
	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).next = node
	}

	//调整尾部位置，及链表元素数量
	(*list).tail = node //node成为新的尾部
	(*list).size++

	return true
}

//删除某一个元素； 输入删除的位置，该位置的节点容器
func (list *List) Remove(i uint64, node *Node2) bool {
	if i >= (*list).size {
		return false
	}

	if i == 0 { //若是首位
		node = (list).head
		(*list).head = (*node).next
		//对头部节点的下一个节点操作； next

		if (*list).size == 1 { //如果只有一个元素，尾部也要调整
			(*list).tail = nil

		}
	} else {
		preItem := (*list).head //定义变量，初始化
		for j := 1; uint64(j) < i; j++ {
			preItem = (*preItem).next //多次next到当前位置的值 ，取出到对应位置的元素，指针是指向地址的值
		}
		node = (*preItem).next         //取出元素的下一个元素，存入 参数的容器
		(*preItem).next = (*node).next // 原来位置的下下下个位置 替代 下下个位置
		if i == ((*list).size - 1) {   //若删除的刚好是最后一位，因为list带有0位，所以数量上减1
			(*list).tail = preItem //把尾巴设置为当前值

		}
	}
	(*list).size--
	return true
}

//插入数据
func (list *List) Insert(i uint64, node *Node2) bool {
	//空的节点、索引超出范围和空链表都无法做插入操作
	if node == nil || i > (*list).size || (*list).size == 0 {
		return false
	}

	if i == 0 { //直接排第一，也
		(*node).next = (*list).head
		(*list).head = node
	} else {
		//找到前一个元素
		preItem := (*list).head
		for j := 1; uint64(j) < i; j++ { //数前面i个元素
			preItem = (*preItem).next
		}
		//原来元素放到新元素后面，新元素放到前一个元素后面
		(*node).next = (*preItem).next //if 外面的u作用域就没有preItem了
		(*preItem).next = preItem
	}
	(*list).size++ //写在函数外面的list也没有作用域
	return true
}

//

//获取特定位置元素
func (list *List) Get(i uint64) *Node2 { //返回的是一个节点结构体指针 的值
	if i >= (*list).size { //接收者的 指针值 里的元素
		return nil
	}

	item := (*list).head //初始化一下
	for j := 0; uint64(j) < i; j++ {
		item = (*item).next // 使用指针，就是不怕重名的元素遍历;只要next一次
		// i位置 i+1的值
	}
	return item

}

func main2() {
	var list = List{}
	list.Init() //初始化自定义结构体，list链表
	for i := 1; i < 100; i++ {
		var node = Node2{data: i} //定义节点并赋值100次
		list.Append(&node)
	}

	var node = list.Get(35)
	fmt.Printf("当前节点位置： %+q \n,数据：%d \n", node, node.data)
	var deleteNode = &Node2{}
	var result = list.Remove(35, deleteNode)
	fmt.Printf("删除是否成功%+q \n", result)

	var node2 = list.Get(35)
	fmt.Printf("当前节点位置： %+q \n，数据：%d\n", node2, node2.data)
	//'+' 总是输出数值的正负号；对%q(%+q)会生成全部是ASCII字符的输出（通过转义）
}

/*
链表本身是一结构体，该结构体内部的元素也是结构体；
属性：
头节点 head
尾节点 tail
元素的容量 size

节点一结构体
属性：
本身数据 data
下一个节点 next


jeremy:
结构体嵌套结构体有些慌
回调自己的类型定义自己也有些慌

链表特点--
利用节点的属性，自己指向遍历下一个
已知链的某位置，可以从元素节点的next元素提取下一个位置的元素节点

*/

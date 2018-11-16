/*
定义了数据类型：链表，并给出相关方法
链表类型： Node
链表的元素类型：Element
*/
package algorithmStudy

import (
	"fmt"
	//"os"
)

//
const (
	ERROR = -10000000001
)

//即使 重命令了Element 为int型，也不能当作int使用！但是方便修改元素的类型；其他代码还是统一一下吧；数据层面代码不修改
type Element int

type Node struct {
	Data Element //
	Next *Node   //
}

//
type Noder interface {
	Add(head *Node, new *Node)                  //
	Delete(head *Node, index int)               //
	Insert(head *Node, index int, data Element) //
	GetLength(head *Node) int                   //
	Search(head *Node, data Element)            //
	GetData(head *Node, index int) Element      //
}

//添加头节点，数据
func Add(head *Node, data Element) {
	point := head //
	for point.Next != nil {
		point = point.Next //移位
	}
	var node Node      //
	point.Next = &node //
	node.Data = data

	head.Data = Element(GetLength(head)) //

	if GetLength(head) > 1 {
		Traverse(head)
	}
}

//删除 头节点 index 位置
func Delete(head *Node, index int) Element {
	//判断index的合法性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return ERROR
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Next //移位
		}
		point.Next = point.Next.Next //赋值
		data := point.Next.Data
		return data
	}
}

//插入头节点 index位置data元素
func Insert(head *Node, index int, data Element) {
	//检验index的合法性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")

	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Next //移位
		}
		var node Node //新节点，赋值
		node.Data = data
		node.Next = point.Next
		point.Next = &node
	}
}

//获取长度头节点
func GetLength(head *Node) int {
	point := head
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

//搜索头结点 data 元素
func Search(head *Node, data Element) {
	point := head
	index := 0
	for point.Next != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Next
			if index > GetLength(head)-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}

//获取Data头结点index位置
func GetData(head *Node, index int) Element {
	point := head
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return ERROR
	} else {
		for i := 0; i < index; i++ {
			point = point.Next
		}
		return point.Data
	}
}

//遍历头节点
func Traverse(head *Node) {
	point := head.Next
	for point.Next != nil {
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("Traverse OK!")
}

//递归单链反转 单链头节点 反转后链表新头节点
func ReverseList(head, node *Node) *Node {
	if node.Next == nil {
		head.Next = node
		return node
	}
	/*
		判断 node移动，是否到链表最后点；
		直到移动到末尾 返回末节点node

	*/

	n := ReverseList(head, node.Next) //递归一遍，原列表节点移动，返回值存入临时指针

	if n != nil { //链表末节点
		n.Next = node // 临时指针移动，指向
		node.Next = nil
	}
	return node

}

//遍历单链反转 三指针 node prev next ，使得node prev连个节点间的指向反向，同时用next记录剩下的链表.
//复杂度 O（n）
func ReverseListV2(head *Node) *Node {

	pNode := head.Next
	head.Next = nil //首节点

	pPrev := head

	pFlo := &Node{}

	for pNode != nil { //判断当前要处理的指针指向的节点不是移动到最后的空节点
		pFlo = pNode.Next  // 后指针 根据当前节点的属性Next,指向后一个节点地址，位移下一节点
		pNode.Next = pPrev //当前节点属性Next改变，指向前指针 的地址
		pPrev = pNode      //前指针 指向 它的下一个节点（即当前操作的节点）指针的地址，位移下一节点
		pNode = pFlo       //当前节点指针 指向 它的下一个节点即后指针的地址，位移下一节点(准备更改下个节点的next属性)

	}
	return pNode

}

//思路2单个链表内部操作，复杂度，主要是熟练在带有链的属性的数据结构内自由调度组合位置

//删除链表中倒数第X位置的节点

/*func main() {
	var head LinkNode = LinkNode{Data: 0, Next: nil}
	head.Data = 0
	var nodeArray []Element
	for i := 0; i < 10; i++ {
		nodeArray = append(nodeArray, Element(i+1+i*100))
		Add(&head, nodeArray[i])
	}

	Delete(&head, 3)
	Search(&head, 2032)
	Insert(&head, 23, 10010)
	Traverse(&head)
	fmt.Println("data is ", GetData(&head, 6))
	fmt.Println("length:", GetLength(&head))
	os.Exit(0)
}
*/

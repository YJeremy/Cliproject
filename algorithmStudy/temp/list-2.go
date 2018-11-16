package algorithmStudy

import (
	"fmt"
	//"os"
)

//
const (
	ERROR = -10000000001
)

type Element int64

type LinkNode struct {
	Data Element   //
	Next *LinkNode //
}

//
type LinkNoder interface {
	Add(head *LinkNode, new *LinkNode)              //
	Delete(head *LinkNode, index int)               //
	Insert(head *LinkNode, index int, data Element) //
	GetLength(head *LinkNode) int                   //
	Search(head *LinkNode, data Element)            //
	GetData(head *LinkNode, index int) Element      //
}

//添加头节点，数据
func Add(head *LinkNode, data Element) {
	point := head //
	for point.Next != nil {
		point = point.Next //移位
	}
	var node LinkNode  //
	point.Next = &node //
	node.Data = data

	head.Data = Element(GetLength(head)) //

	if GetLength(head) > 1 {
		Traverse(head)
	}
}

//删除 头节点 index 位置
func Delete(head *LinkNode, index int) Element {
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
func Insert(head *LinkNode, index int, data Element) {
	//检验index的合法性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")

	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Next //移位
		}
		var node LinkNode //新节点，赋值
		node.Data = data
		node.Next = point.Next
		point.Next = &node
	}
}

//获取长度头节点
func GetLength(head *LinkNode) int {
	point := head
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

//搜索头结点 data 元素
func Search(head *LinkNode, data Element) {
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
func GetData(head *LinkNode, index int) Element {
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
func Traverse(head *LinkNode) {
	point := head.Next
	for point.Next != nil {
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("Traverse OK!")
}

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

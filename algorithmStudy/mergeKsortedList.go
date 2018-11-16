//把K个链表合并并排序

package algorithmStudy

//思路2：遍历合并
func MergeKLists2(lists []*Node) *Node {
	//特殊链表检查
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	//检查
	var nodeValues []Element
	for _, list := range lists {
		for list != nil {
			nodeValues = append(nodeValues, list.Data)
			list = list.Next //链表移动
		}

	}
	insertSort(nodeValues)
	curNode := &Node{}
	dummy := curNode
	for _, val := range nodeValues { //创建列表
		curNode.Next = &Node{val, nil}
		curNode = curNode.Next
	}
	return dummy.Next

}

//插入排序
func insertSort(arr []Element) []Element {
	len := len(arr)
	if len < 2 {
		return arr
	}
	for i := 1; i < len; i++ {
		for j := i - 1; j >= 0; j-- {
			//前面的值大则交换
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

//思路3：分治法；即遍历数组前和中间两两合并，确保所有的列表长度差不大，遍历次数均匀
func MergeKLists3(lists []*Node) *Node {

	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	step := len(lists)

	for step > 1 {
		//找到中间位置
		mid := (step + 1) / 2
		//计算前半部分的合并结果
		//注意mid 和i的上限取值差1,保证mid在i的后边
		for i := 0; i < step/2; i++ {
			lists[i] = MergeTwoLists(lists[i], lists[i+mid])
		}
		//时间上的一次处理，空间上的多次处理；

		step = mid
	}
	return lists[0]

}

func MergeTwoLists(l1 *Node, l2 *Node) *Node {
	//特殊链表处理
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	cur1 := l1
	cur2 := l2
	res := &Node{0, nil}
	dummy := res
	for cur1 != nil && cur2 != nil {
		res.Next = &Node{}
		res = res.Next
		if cur1.Data > cur2.Data {
			res.Data = cur2.Data
			cur2 = cur2.Next
		} else {
			res.Data = cur1.Data
			cur1 = cur1.Next
		}
	}
	if cur1 != nil {
		res.Next = cur1
	}
	if cur2 != nil {
		res.Next = cur2
	}
	return dummy.Next

}

package main

import "fmt"

/**
	反转单链表
 */
//定义一个结构体
type node struct {
	value    int
	nextNode *node
}

func reverseNode(head *node) *node {
	//先申明两个变量
	//前一个节点
	var preNode *node
	preNode = nil
	//后一个节点
	nextNode := new(node)
	nextNode = nil
	for head != nil {
		//保存头节点的下一个节点
		nextNode = head.nextNode
		//将头节点指向前一个节点
		head.nextNode = preNode
		//更新前一个节点
		preNode = head
		//更新头节点
		head = nextNode
	}
	return preNode
}

//这里是打印一个节点
func printNode(head *node) {
	if head != nil {
		//fmt.Println(head)
		fmt.Printf("%s\t", head)
		head = head.nextNode
	}
	fmt.Println()
}
func main() {
	node1 := new(node)
	node1.value = 1
	node2 := new(node)
	node2.value = 2
	node3 := new(node)
	node3.value = 3
	node4 := new(node)
	node4.value = 4
	node1.nextNode = node2
	node2.nextNode = node3
	node3.nextNode = node4
	printNode(node1)
	head := reverseNode(node1)
	printNode(head)

}

package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func (lnode *ListNode) printList() {
	nextNode := lnode
	for nextNode != nil {
		fmt.Print(nextNode.val)
		if nextNode.next != nil {
			fmt.Print(" ")
		}
		nextNode = nextNode.next
	}
}

func (lnode *ListNode) reverseIteratively(headNode *ListNode) {
	prevNode := headNode
	currNode := headNode.next
	for currNode != nil {
		tmpNode := currNode.next
		currNode.next = prevNode
		prevNode = currNode
		currNode = tmpNode
	}
	headNode.next = nil
}

func (lnode *ListNode) reverseRecursively(node *ListNode) *ListNode {
	if node.next != nil {
		retNode := lnode.reverseRecursively(node.next)
		retNode.next = node
		node.next = nil
	}
	return node
}

func main() {
	fmt.Println("Hello World")
	// Test Program
	// Initialize the test list:
	testHead := ListNode{val: 4}
	node1 := ListNode{val: 3}
	testHead.next = &node1
	node2 := ListNode{val: 2}
	node1.next = &node2
	node3 := ListNode{val: 1}
	node2.next = &node3
	testTail := ListNode{val: 0}
	node3.next = &testTail

	print("Initial list: ")
	testHead.printList()
	testHead.reverseIteratively(&testHead)
	// testHead.reverseRecursively(&testHead)
	print("\nList after reversal: ")
	testTail.printList()
	print("\n")
}

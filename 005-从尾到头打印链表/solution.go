package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

var arr []int

func printListFromTailToHead(head ListNode) []int {
	if head.val == 0 {
		return arr
	}

	if head.val != 0 {
		if head.next != nil {
			printListFromTailToHead(*head.next)
		}
		arr = append(arr, head.val)
	}
	return arr

}

func main() {
	a := ListNode{1, nil}
	b := ListNode{2, &a}
	c := ListNode{3, &b}
	d := ListNode{4, &c}
	e := ListNode{5, &d}
	//fmt.Printf("%+v", e)
	fmt.Println(printListFromTailToHead(e))
}

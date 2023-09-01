package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 4}

	fmt.Println(*head)
	fmt.Println(&head)
	fmt.Println(*head.Next)
	fmt.Println(&head.Next)
	fmt.Println(*head.Next.Next)
	fmt.Println(&head.Next.Next)
	fmt.Println(*head.Next.Next.Next)
	fmt.Println(&head.Next.Next.Next)
	fmt.Println(head)

	head = head.Next
	curr := head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil

	curr = head
	for curr != nil {
		fmt.Printf("%d", curr.Value)
		curr = curr.Next
	}

}

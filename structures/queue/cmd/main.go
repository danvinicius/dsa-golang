package main

import (
	"fmt"

	queue "github.com/danvinicius/dsa-golang/structures/queue/pkg"
)

func main() {
	queue := queue.NewQueue()

	queue.Enqueue(5)
	fmt.Println(queue.Items())

	queue.Enqueue(4)
	fmt.Println(queue.Items())

	queue.Enqueue(7)
	fmt.Println(queue.Items())

	queue.Enqueue(11)
	fmt.Println(queue.Items())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Items())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Items())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Items())

	queue.Enqueue(67)
	fmt.Println(queue.Items())
	
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Items())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Items())
}

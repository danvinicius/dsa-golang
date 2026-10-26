package main

import (
	"fmt"

	stack "github.com/danvinicius/dsa-golang/structures/stack/pkg"
)

func main() {
	stack := stack.NewStack()

	stack.Push(5)
	fmt.Println(stack.Items())

	stack.Push(4)
	fmt.Println(stack.Items())

	stack.Push(7)
	fmt.Println(stack.Items())

	stack.Push(11)
	fmt.Println(stack.Items())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Items())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Items())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Items())

	stack.Push(67)
	fmt.Println(stack.Items())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Items())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Items())
}

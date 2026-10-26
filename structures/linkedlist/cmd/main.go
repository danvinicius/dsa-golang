package main

import (
	"fmt"

	linked_list "github.com/danvinicius/dsa-golang/structures/linked-list/pkg"
)

func main() {
	linkedList := linked_list.NewLinkedList()

	linkedList.InsertFront(5)
	linkedList.InsertEnd(67)
	linkedList.InsertEnd(69)
	linkedList.InsertFront(12)

	linkedList.Delete(69)

	fmt.Println(linkedList.Travesal())
}

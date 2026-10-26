package binarysearchtree

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

type BinarySearchTree[T cmp.Ordered] struct {
	Root *Node[T]
}

func NewNode[T cmp.Ordered](data T) *Node[T] {
	return &Node[T]{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func NewBinarySearchTree[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) Insert(value T) {
	var parent *Node[T]
	x := bst.Root

	for x != nil {
		parent = x

		if value < x.Data {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	if parent == nil {
		bst.Root = NewNode(value)
	} else if value < parent.Data {
		parent.Left = NewNode(value)
	} else {
		parent.Right = NewNode(value)
	}
}

func (bst *BinarySearchTree[T]) Search(value T) (*Node[T], bool) {
	x := bst.Root

	for x != nil {
		if x.Data == value {
			return x, true
		}
		if value < x.Data {
			x = x.Left
		}
		if value > x.Data {
			x = x.Right
		}
	}
	return nil, false
}

func (t *BinarySearchTree[T]) Height(node *Node[T]) int {
	if node == nil {
		node = t.Root
	}

	hLeft := 0
	hRight := 0

	if node.Left != nil {
		hLeft = t.Height(node.Left)
	}
	if node.Right != nil {
		hRight = t.Height(node.Right)
	}

	if hLeft > hRight {
		return hLeft + 1
	}
	return hRight + 1
}

func (t *BinarySearchTree[T]) invert(node *Node[T]) {
	if node != nil {
		node.Left, node.Right = node.Right, node.Left
		t.invert(node.Left)
		t.invert(node.Right)
	}
}

func (t *BinarySearchTree[T]) Invert() {
	t.invert(t.Root)
}

// In Order (simetric order): left -> root -> right
func (t *BinarySearchTree[T]) InOrderTraversal(node *Node[T]) {
	if node == nil {
		node = t.Root
	}
	if node.Left != nil {
		t.InOrderTraversal(node.Left)
	}
	fmt.Print(node.Data, ",")
	if node.Right != nil {
		t.InOrderTraversal(node.Right)
	}
}

// Post Order: left -> right -> root
func (t *BinarySearchTree[T]) PostOrderTraversal(node *Node[T]) {
	if node == nil {
		node = t.Root
	}
	if node.Left != nil {
		t.PostOrderTraversal(node.Left)
	}
	if node.Right != nil {
		t.PostOrderTraversal(node.Right)
	}
	fmt.Print(node.Data, ",")
}

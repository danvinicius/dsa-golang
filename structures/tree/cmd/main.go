package main

import (
	"fmt"

	binary_search_tree "github.com/danvinicius/dsa-golang/structures/tree/pkg/binary-search-tree"
)

/**
in order travesal

		'+'
	 /    \
'a'	   '*'
      /   \
		'b'   '-'
				 /   \
			 '/'   'e'
			/  \
		'c'  'd'

*/

/**
post order travesal

		'l'
	 /    \
'd'	   'e'
			/   \
	 'n'   'i'
	/
'a'

*/

func main() {

	/**
	  		 10
	  	 /   \
	    5     15
		/	 \   /  \
	 0    8 11   19

	 **/

	bst := binary_search_tree.NewBinarySearchTree[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(0)
	bst.Insert(8)
	bst.Insert(15)
	bst.Insert(11)
	bst.Insert(19)
	fmt.Println("Altura:", bst.Height(nil))

	fmt.Println(bst.Search(15))

	bst.InOrderTraversal(nil)
	fmt.Println()

	bst.Invert()

	bst.InOrderTraversal(nil)
	fmt.Println()
}

package main

import (
	"fmt"

	hash_table "github.com/danvinicius/dsa-golang/structures/hash-table/pkg"
	"github.com/danvinicius/dsa-golang/structures/hash-table/pkg/hasher"
)

func main() {
	badHasher := hasher.NewBadHasher()
	hashTable := hash_table.NewHashTable(100, badHasher)
	hashTable.Put("tralalero tralala", 25)
	hashTable.Put("tung tung tung sahur", 67)
	fmt.Println(hashTable.Get("tung tung tung sahur"))
}

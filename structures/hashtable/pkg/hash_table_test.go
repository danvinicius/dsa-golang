package hashtable

import (
	"fmt"
	"testing"

	"github.com/danvinicius/dsa-golang/structures/hash-table/pkg/hasher"
)

func prepareHashTableWithBadHasher(size int) *HashTable {
	badHasher := hasher.NewBadHasher()
	return NewHashTable(size, badHasher)
}

func TestResize(t *testing.T) {
	hashTable := prepareHashTableWithBadHasher(4)

	hashTable.Put("ab", 99)
	hashTable.Put("ba", 88)
	for i := range 50 {
		hashTable.Put(fmt.Sprintf("chave-%d", i), i)
	}

	t.Run("Elements with same hash value still exist after resize", func(t *testing.T) {
		chaveAB, _ := hashTable.Get("ab")
		chaveBA, _ := hashTable.Get("ba")

		expectedChaveABValue := 99
		expectedChaveBAValue := 88

		if chaveAB != expectedChaveABValue || chaveBA != expectedChaveBAValue {
			t.Error("Values for keys 'ab' and 'ba' do not match expected values.")
		}
	})

	t.Run("All elements still exist", func(t *testing.T) {
		for i := range 50 {
			if v, ok := hashTable.Get(fmt.Sprintf("chave-%d", i)); !ok || v != i {
				t.Errorf("Value for key 'chave-%d' does not match expected value.", i)
			}
		}
	})
}

func TestDelete(t *testing.T) {
	hashTable := prepareHashTableWithBadHasher(8)

	t.Run("Delete non-existent key", func(t *testing.T) {

		expectedCount := hashTable.Count()
		hashTable.Delete("non-existent key")

		if expectedCount != hashTable.Count() {
			t.Errorf("New count %d is differente from expected count %d.", hashTable.Count(), expectedCount)
		}
	})
}

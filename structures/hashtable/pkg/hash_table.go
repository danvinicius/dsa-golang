package hashtable

const (
	defaultSize       = 8
	maxLoadFactor     = 0.75
	growthThreshold   = 1000
	smallGrowthFactor = 2.0
	largeGrowthFactor = 1.5
)

type Hasher interface {
	Hash(key string, size int) int
}

type entry struct {
	key   string
	value any
}

type HashTable struct {
	count   int
	buckets [][]entry
	hasher  Hasher
}

func NewHashTable(size int, hasher Hasher) *HashTable {
	if size < 1 {
		size = defaultSize
	}
	return &HashTable{
		count:   0,
		hasher:  hasher,
		buckets: make([][]entry, size),
	}
}

func (t *HashTable) Len() int {
	return len(t.buckets)
}

func (t *HashTable) Count() int {
	return t.count
}

func (t *HashTable) Get(key string) (any, bool) {
	index := t.hasher.Hash(key, t.Len())
	for _, e := range t.buckets[index] {
		if e.key == key {
			return e.value, true
		}
	}
	return nil, false
}

func (t *HashTable) Put(key string, value any) {
	index := t.hasher.Hash(key, t.Len())
	for i := range t.buckets[index] {
		if t.buckets[index][i].key == key {
			t.buckets[index][i].value = value
			return
		}
	}
	t.count++
	t.buckets[index] = append(t.buckets[index], entry{key: key, value: value})

	if t.LoadFactor() > maxLoadFactor {
		t.Resize()
	}
}

func (t *HashTable) Delete(key string) {
	index := t.hasher.Hash(key, t.Len())
	last := len(t.buckets[index]) - 1
	for i := range t.buckets[index] {
		if t.buckets[index][i].key == key {
			t.buckets[index][i] = t.buckets[index][last]
			t.buckets[index][last] = entry{}
			t.buckets[index] = t.buckets[index][:last]
			t.count--
			return
		}
	}
}

func (t *HashTable) LoadFactor() float64 {
	return float64(t.Count()) / float64(t.Len())
}

func (t *HashTable) Resize() {
	newSize := int(float64(t.Len()) * smallGrowthFactor)
	if t.Len() > growthThreshold {
		newSize = int(float64(t.Len()) * largeGrowthFactor)
	}
	newBuckets := make([][]entry, newSize)
	for i := range t.buckets {
		for _, e := range t.buckets[i] {
			index := t.hasher.Hash(e.key, newSize)
			newBuckets[index] = append(newBuckets[index], e)
		}
	}
	t.buckets = newBuckets
}

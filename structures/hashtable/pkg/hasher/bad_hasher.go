package hasher

type BadHasher struct{}

func (BadHasher) Hash(key string, size int) int {
	var sum = 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % size
}

func NewBadHasher() *BadHasher {
	return &BadHasher{}
}

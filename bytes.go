package levenshtein

import "math"

type bytesFrom struct {
	indelCost int
	subCost   int
	from      []byte
	costs     []int
}

// Dist calculates levenshtein distance between 'from' and 'to'.
// If actual distance is less or equal then maxCost,
// then function returns accurate value of it.
// But otherwise returns value which is greater than maxCost.
func (b *bytesFrom) Dist(to []byte, maxCost int) int {
	l := len(b.from)
	m := b.costs
	for i := 1; i <= l; i++ {
		m[i] = i * b.indelCost
	}
	lastdiag := 0
	for x, rx := range to {
		m[0], lastdiag = (x+1)*b.indelCost, x*b.indelCost
		min := math.MaxInt32
		for y, ry := range b.from {
			m[y+1], lastdiag = min3(m[y]+b.indelCost, m[y+1]+b.indelCost, lastdiag+b.mkSubCost(rx, ry)), m[y+1]
			if m[y+1] < min {
				min = m[y+1]
			}
		}
		if min > maxCost {
			return min
		}
	}
	return m[l]
}

func (b *bytesFrom) mkSubCost(b1, b2 byte) int {
	if b1 == b2 {
		return 0
	}
	return b.subCost
}

// BytesFrom computes distance from one byte array to other byte arrays.
// Instances should not be used concurrently.
type BytesFrom interface {
	Dist(to []byte, maxCost int) int
}

// FromBytes return BytesFrom for a given bytes array
func FromBytes(from []byte) BytesFrom {
	l := len(from)
	m := make([]int, l+1)
	return &bytesFrom{
		indelCost: 1,
		subCost:   1,
		from:      from,
		costs:     m,
	}
}

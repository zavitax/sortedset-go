package sortedset

import "golang.org/x/exp/constraints"

type SortedSetLevel[K constraints.Ordered, SCORE constraints.Ordered, V any] struct {
	forward *SortedSetNode[K, SCORE, V]
	span    int64
}

// Node in skip list
type SortedSetNode[K constraints.Ordered, SCORE constraints.Ordered, V any] struct {
	key      K     // unique key of this node
	Value    V     // associated data
	score    SCORE // score to determine the order of this node in the set
	backward *SortedSetNode[K, SCORE, V]
	level    []SortedSetLevel[K, SCORE, V]
}

// Get the key of the node
func (this *SortedSetNode[K, SCORE, V]) Key() K {
	return this.key
}

// Get the node of the node
func (this *SortedSetNode[K, SCORE, V]) Score() SCORE {
	return this.score
}

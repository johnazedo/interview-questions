package datastructures

/*
DESCRIPTION:  A priority queue is a type of queue that arranges elements based on their priority values.
Elements with higher priority values are typically retrieved before elements with lower priority values.

OBS: Append built-in function has O(N) time complexity in the worst case, so avoid it
*/

// TODO: Create a priority queue using heap

type Item struct {
	Priority int
	Value    any
}

type PQ struct {
	size     int
	capacity int
	Queue    []Item
}

func NewPQ(capacity int) *PQ {
	return &PQ{
		size:     0,
		capacity: capacity,
		Queue:    make([]Item, capacity, capacity),
	}
}

func (pq *PQ) Add(value Item) {
	// Can add anywhere
	// Time: O(1)
	if !pq.Full() {
		pq.Queue[pq.size] = value
		pq.size++
	}
}

func (pq *PQ) GetMax() Item {
	// It's necessary iterate through the array to get the highest priority item
	// Time: O(N)
	max := 0
	index := 0
	for i, v := range pq.Queue {
		if v.Priority > max {
			max = v.Priority
			index = i
		}
	}

	item := pq.Queue[index]
	pq.swap(index)

	// Remove the last element
	pq.Queue = pq.Queue[:pq.size-1]
	pq.size--
	return item
}

func (pq *PQ) swap(index int) {
	// Puts the element in the last position and puts the last element in the index position
	pq.Queue[index], pq.Queue[pq.size-1] = pq.Queue[pq.size-1], pq.Queue[index]
}

func (pq *PQ) GetSize() int {
	return pq.size
}

func (pq *PQ) Empty() bool {
	return pq.size == 0
}

func (pq *PQ) Full() bool {
	return pq.size == pq.capacity
}

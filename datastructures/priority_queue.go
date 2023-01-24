package datastructures

/*
DESCRIPTION:  A priority queue is a type of queue that arranges elements based on their priority values.
Elements with higher priority values are typically retrieved before elements with lower priority values.

OBS: Append built-in function has O(N) time complexity in the worst case, so avoid it
*/

// TODO: Create a priority queue using heap
// TODO: Create test to this implementation

type Item struct {
	Priority int
	Value    any
}

type PQ struct {
	Size     int
	Capacity int
	Queue    []Item
}

func NewPQ(capacity int) *PQ {
	return &PQ{
		Size:     0,
		Capacity: capacity,
		Queue:    make([]Item, capacity, capacity),
	}
}

func (pq *PQ) Add(value Item) {
	// Can add anywhere
	// Time: O(1)
	if !pq.Full() {
		pq.Queue[pq.Size] = value
		pq.Size++
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
	pq.Queue = pq.Queue[:pq.Size-1]
	pq.Size--
	return item
}

func (pq *PQ) swap(index int) {
	// Puts the element in the last position and puts the last element in the index position
	pq.Queue[index], pq.Queue[pq.Size-1] = pq.Queue[pq.Size-1], pq.Queue[index]
}

func (pq *PQ) GetSize() int {
	return pq.Size
}

func (pq *PQ) Empty() bool {
	return pq.Size == 0
}

func (pq *PQ) Full() bool {
	return pq.Size == pq.Capacity
}

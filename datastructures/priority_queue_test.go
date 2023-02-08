package datastructures

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPQ(3)
	expected := Item{3, 3}
	pq.Add(Item{1, 1})
	pq.Add(Item{2, 2})
	pq.Add(Item{3, 3})

	// Test Add()
	if len(pq.Queue) != 3 {
		t.Fatalf("Add() isn't working")
	}

	if pq.GetSize() != len(pq.Queue) {
		t.Fatalf("Add() isn't counting the number of items correctly")
	}

	// Test GetMax()
	result := pq.GetMax()

	if result != expected {
		t.Fatalf("GetMax() can't find the maximun priority")
	}

	if pq.GetSize() != 2 {
		t.Fatalf("GetMax() isn't removing the max priority item")
	}
}

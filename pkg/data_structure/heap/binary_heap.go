package heap

import (
	"fmt"
	ADT "github.com/nathan-tw/algo/pkg/abstact_data_type"
)

type HeapNode map[int]int

func (hn HeapNode) Min(a, b int) (index int) {
	if hn[a] > hn[b] {
		return b
	}
	return a
}

type Heap struct {
	ADT.PriorityQueue
	Nodes HeapNode
}

func NewBinaryHeapWithElem(elements []int) *Heap {
	hp := new(Heap)
	m := make(HeapNode)
	for index, ele := range elements {
		m[index+1] = ele
	}
	hp.Nodes = m
	return hp
}

func NewBinaryHeap() *Heap {
	hp := new(Heap)
	return hp
}

func (hp *Heap) MinHeapify(key int) {
	minKey := key
	for key*2 <= len(hp.Nodes) {
		// if only left child
		if _, ok := hp.Nodes[key*2+1]; !ok {
			// swap
			minKey = hp.Nodes.Min(key, key*2)
			if minKey != key {
				hp.Nodes[key], hp.Nodes[minKey] = hp.Nodes[minKey], hp.Nodes[key]
				return
			} else {
				return
			}
		} else {
			minKey = hp.Nodes.Min(key*2, key*2+1)
			minKey = hp.Nodes.Min(key, minKey)
			if minKey != key {
				hp.Nodes[key], hp.Nodes[minKey] = hp.Nodes[minKey], hp.Nodes[key]
				key = minKey
			} else {
				return
			}
		}
	}
}

func (hp *Heap) BuildMinHeap() {
	for key := len(hp.Nodes); key >= 1; key-- {
		hp.MinHeapify(key)
	}
}

func (hp *Heap) Insert(input int) {
	maxKey := len(hp.Nodes) + 1
	hp.Nodes[maxKey] = input
	for maxKey != 1 {		
		if hp.Nodes[maxKey] < hp.Nodes[maxKey/2] {
			hp.Nodes[maxKey], hp.Nodes[maxKey/2] = hp.Nodes[maxKey/2], hp.Nodes[maxKey]
			maxKey /= 2
		} else{
			return
		}
	} 	
}

func (hp *Heap) GetMinimum() int {
	return hp.Nodes[1]
}

func (hp *Heap) ExtractMin() int {
	min := hp.Nodes[1]
	length := len(hp.Nodes)
	hp.Nodes[1] = hp.Nodes[length]
	delete(hp.Nodes, length)
	hp.MinHeapify(1)

	return min
}

func (hp *Heap) IncreaseKey(index int, incr int) {
	hp.Nodes[index] += incr
	hp.MinHeapify(index)
}

func (hp *Heap) PrintNodes() {
	for k, v := range hp.Nodes {
		fmt.Println(k,v)
	}
	fmt.Println()
}

package heap

import (
	"fmt"

	ADT "github.com/nathan-tw/algo/pkg/abstact_data_type"
)

type Heap struct {
	ADT.PriorityQueue
	Nodes HeapNode
}

type HeapNode []int

func (hn HeapNode) min(i1, i2 int) (index int) {
	if hn[i1] > hn[i2] {
		return i2
	}
	return i1
}

func NewBinaryHeap(elements []int) *Heap {
	nodes := make([]int, 1)
	nodes = append(nodes, elements...)
	return &Heap{
		Nodes: nodes,
	}
}

func (hp *Heap) PrintHeap() {
	for i := 1; i < len(hp.Nodes); i++ {
		fmt.Println(hp.Nodes[i])
	}
}

func (hp *Heap) minHeapify(index int) {
	minIndex := index
	for index*2 <= len(hp.Nodes)-1 {
		// if only left node exist
		if index*2+1 > len(hp.Nodes)-1 {
			minIndex = hp.Nodes.min(index, index*2)
			if minIndex != index {
				// swap
				hp.Nodes[index], hp.Nodes[minIndex] = hp.Nodes[minIndex], hp.Nodes[index]
				return
			} else {
				return
			}
		} else {
			minIndex = hp.Nodes.min(index*2, index*2+1)
			minIndex = hp.Nodes.min(index, minIndex)
			if minIndex != index {
				hp.Nodes[index], hp.Nodes[minIndex] = hp.Nodes[minIndex], hp.Nodes[index]
				index = minIndex
			} else {
				return
			}
		}
	}
}

func (hp *Heap) BuildMinHeap() {
	for index := len(hp.Nodes)-1; index >= 1; index-- {
		hp.minHeapify(index)
	}
}

func (hp *Heap) GetMinimum() int {
	return hp.Nodes[1]
}

func (hp *Heap) Insert(input int) {
	maxIndex := len(hp.Nodes)
	hp.Nodes = append(hp.Nodes, input)
	for maxIndex != 1 {		
		if hp.Nodes[maxIndex] < hp.Nodes[maxIndex/2] {
			hp.Nodes[maxIndex], hp.Nodes[maxIndex/2] = hp.Nodes[maxIndex/2], hp.Nodes[maxIndex]
			maxIndex /= 2
		} else{
			return
		}
	} 	
}

func (hp *Heap) ExtractMin() int {
	// replace the min with the least index element
	minNode := hp.Nodes[1]
	last := len(hp.Nodes)-1
	hp.Nodes[1] = hp.Nodes[last]
	hp.Nodes = hp.Nodes[:last]
	return minNode
}

func (hp *Heap) IncreaseKey(index int, incr int) {
	hp.Nodes[index] += incr
	hp.minHeapify(index)
}
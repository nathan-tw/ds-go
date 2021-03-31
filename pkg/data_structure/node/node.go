package node

type BinaryNode struct {
	Value int
	Left *BinaryNode
	Right *BinaryNode
}

type SingleNode struct {
	Value int
	Next *SingleNode
}

type DoubleNode struct {
	Value int
	Next *DoubleNode
	Prev *DoubleNode
}
package binary_tree

import "github.com/nathan-tw/algo/pkg/data_structure/node"


type AVLTree node.BinaryNode


func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

package model

type BinaryTree struct {
	Nodes map[string]*Node `json:"nodes"`
	Root  string           `json:"root"`
}

type Tree struct {
	Tree *BinaryTree `json:"tree"`
}

type Node struct {
	ID    string  `json:"id"`
	Left  *string `json:"left,omitempty"`
	Right *string `json:"right,omitempty"`
	Value int     `json:"value"`
}

type BinaryTreeSum struct {
	Sum int `json:"sum"`
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

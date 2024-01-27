package model

type BinaryTree struct {
	Nodes []*Node `json:"nodes"`
	Root  string  `json:"root"`
}

type Tree struct {
	Tree *BinaryTree `json:"tree"`
}

type Node struct {
	ID    string  `json:"id"`
	Left  *string `json:"left"`
	Right *string `json:"right"`
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

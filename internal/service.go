package internal

import "main/model"

type Service struct {
}

type Actions interface {
	PostBinaryTreeSum(tree *model.Tree) (*model.BinaryTreeSum, error)
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) PostBinaryTreeSum(tree *model.Tree) (*model.BinaryTreeSum, error) {

	binaryTree := tree.Tree

	root := convertToTreeNode(binaryTree.Nodes[binaryTree.Root], binaryTree.Nodes)
	result := maxPathSum(root)

	return &model.BinaryTreeSum{Sum: result}, nil
}

func convertToTreeNode(node *model.Node, nodes map[string]*model.Node) *model.TreeNode {
	if node == nil {
		return nil
	}

	var leftNode, rightNode *model.Node

	if node.Left != nil {
		leftNode = nodes[*node.Left]
	}
	if node.Right != nil {
		rightNode = nodes[*node.Right]
	}

	return &model.TreeNode{
		Val:   node.Value,
		Left:  convertToTreeNode(leftNode, nodes),
		Right: convertToTreeNode(rightNode, nodes),
	}
	//if node == nil {
	//	return nil
	//}
	//return &model.TreeNode{
	//	Val:   node.Value,
	//	Left:  convertToTreeNode(node.Left),
	//	Right: convertToTreeNode(node.Right),
	//}
}

func maxPathSum(root *model.TreeNode) int {
	maxPathSum := root.Val
	dfs(root, &maxPathSum)
	return maxPathSum
}

func dfs(root *model.TreeNode, maxPathSum *int) int {
	if root == nil {
		return 0
	}

	leftMax := max(dfs(root.Left, maxPathSum), 0)
	rightMax := max(dfs(root.Right, maxPathSum), 0)

	*maxPathSum = max(*maxPathSum, root.Val+leftMax+rightMax)
	return root.Val + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

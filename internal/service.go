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

	nodesMap := make(map[string]*model.Node)
	for _, node := range binaryTree.Nodes {
		nodesMap[node.ID] = node
	}

	root := convertToTreeNode(nodesMap[binaryTree.Root], nodesMap)
	result := maxPathSum(root)

	return &model.BinaryTreeSum{Sum: result}, nil
}

func convertToTreeNode(node *model.Node, nodesMap map[string]*model.Node) *model.TreeNode {
	if node == nil {
		return nil
	}

	var leftNode, rightNode *model.Node

	if node.Left != nil {
		leftNode = nodesMap[*node.Left]
	}

	if node.Right != nil {
		rightNode = nodesMap[*node.Right]
	}

	return &model.TreeNode{
		Val:   node.Value,
		Left:  convertToTreeNode(leftNode, nodesMap),
		Right: convertToTreeNode(rightNode, nodesMap),
	}
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

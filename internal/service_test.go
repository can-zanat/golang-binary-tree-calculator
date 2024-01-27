package internal

import (
	"main/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_PostBinaryTreeSum(t *testing.T) {
	tests := []struct {
		name     string
		tree     *model.Tree
		expected int
	}{
		{
			name:     "Test Case 1",
			tree:     testCaseBody1,
			expected: 18,
		},
		{
			name:     "Test Case 2",
			tree:     testCaseBody2,
			expected: 6,
		},
		{
			name:     "Test Case 3",
			tree:     testCaseBody3,
			expected: 154,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService()
			result, err := s.PostBinaryTreeSum(tt.tree)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, result.Sum)
		})
	}
}

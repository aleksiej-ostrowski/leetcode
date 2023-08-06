package main

import (
	"testing"
)

func Test_1(t *testing.T) {

	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{4, nil, nil}
	tree5 := &TreeNode{5, nil, nil}
	tree2 := &TreeNode{2, tree4, tree5}
	tree1 := &TreeNode{1, tree2, tree3}

	res := diameterOfBinaryTree(tree1)

	if res != 3 {
		t.Fatalf("[1,2,3,4,5] = %d, want %d", res, 3)
	}
}

func Test_2(t *testing.T) {

	tree2 := &TreeNode{2, nil, nil}
	tree1 := &TreeNode{1, tree2, nil}

	res := diameterOfBinaryTree(tree1)

	if res != 1 {
		t.Fatalf("[1,2] = %d, want %d", res, 1)
	}
}

func Test_3(t *testing.T) {

	tree1 := &TreeNode{1, nil, nil}

	res := diameterOfBinaryTree(tree1)

	if res != 0 {
		t.Fatalf("[1] = %d, want %d", res, 0)
	}
}

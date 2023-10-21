package main

import "github.com/xackery/wlk/walk"

type TreeNode struct {
	name     string
	parent   *TreeNode
	children []*TreeNode
}

func NewTreeNode(name string, parent *TreeNode) *TreeNode {
	tn := new(TreeNode)
	tn.name = name
	tn.parent = parent
	return tn
}

// Text returns the name of a tree node
func (tn *TreeNode) Text() string {
	return tn.name
}

func (tn *TreeNode) Parent() walk.TreeItem {
	if tn.parent == nil {
		return nil
	}

	return tn.parent
}

func (tn *TreeNode) ChildCount() int {
	if tn.children == nil {
		return 0
	}
	return len(tn.children)
}

func (tn *TreeNode) ChildAt(index int) walk.TreeItem {
	return tn.children[index]
}

func (tn *TreeNode) Image() interface{} {
	return nil
}

func (tn *TreeNode) ResetChildren() error {
	tn.children = nil

	return nil
}

func (tn *TreeNode) RemoveChild(child walk.TreeItem) {
	childNode, ok := child.(*TreeNode)
	if !ok {
		return
	}

	for i, c := range tn.children {
		if c == childNode {
			tn.children = append(tn.children[:i], tn.children[i+1:]...)
			return
		}
	}
}

func (tn *TreeNode) ChildAdd(name string) *TreeNode {
	child := new(TreeNode)
	child.name = name
	child.parent = tn
	tn.children = append(tn.children, child)
	return child
}

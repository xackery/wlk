package main

import "github.com/xackery/wlk/walk"

type TreeModel struct {
	walk.TreeModelBase
	roots []*TreeNode
}

// NewTreeModel creates a new tree model
func NewTreeModel(name string, parent *TreeNode) *TreeModel {
	tm := new(TreeModel)
	return tm
}

func (tm *TreeModel) LazyPopulation() bool {
	return true
}

func (tm *TreeModel) RootCount() int {
	return len(tm.roots)
}

func (tm *TreeModel) RootAt(index int) walk.TreeItem {
	return tm.roots[index]
}

func (tm *TreeModel) RootAdd(name string) *TreeNode {
	root := new(TreeNode)
	root.name = name
	tm.roots = append(tm.roots, root)
	return root
}

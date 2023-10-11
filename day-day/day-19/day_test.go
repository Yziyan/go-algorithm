// @Author: Ciusyan 10/8/23

package day_19

import (
	"fmt"
	"math/rand"
	"testing"
)

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func TestCodec(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Right.Left = NewTreeNode(4)
	root.Right.Right = NewTreeNode(5)

	codec := Constructor()
	serialize := codec.serialize(root)
	t.Log(serialize)
	root = codec.deserialize(serialize)
	fmt.Println()
}

func TestNTreeNode2TreeNode(t *testing.T) {
	// N 叉树
	//				1
	//		2 		3 			4
	//	5 6 7		8 		9 10 11
	nRoot := &NTreeNode{
		Val: 1,
		children: []*NTreeNode{
			&NTreeNode{
				Val: 2,
				children: []*NTreeNode{
					&NTreeNode{
						Val:      5,
						children: []*NTreeNode{},
					},
					&NTreeNode{
						Val:      6,
						children: []*NTreeNode{},
					},
					&NTreeNode{
						Val:      7,
						children: []*NTreeNode{},
					},
				},
			},
			&NTreeNode{
				Val: 3,
				children: []*NTreeNode{
					&NTreeNode{
						Val:      8,
						children: []*NTreeNode{},
					},
				},
			},
			&NTreeNode{
				Val: 4,
				children: []*NTreeNode{
					&NTreeNode{
						Val:      9,
						children: []*NTreeNode{},
					},
					&NTreeNode{
						Val:      10,
						children: []*NTreeNode{},
					},
					&NTreeNode{
						Val:      11,
						children: []*NTreeNode{},
					},
				},
			},
		},
	}
	// 二叉树
	root := Encode(nRoot)
	nRoot = Decode(root)
	fmt.Println()
}

func RandomTree(level int, maxLevel int, maxVal int) *TreeNode {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}

	// 建根节点
	root := NewTreeNode(1 + rand.Intn(maxVal))
	// 递归建左右子树
	root.Left = RandomTree(level+1, maxLevel, maxVal)
	root.Right = RandomTree(level+1, maxLevel, maxVal)
	return root
}

func TestMaxWidth(t *testing.T) {
	root := RandomTree(0, 5, 20)

	width := MaxWidth(root)
	width2 := MaxWidth1(root)

	t.Log(width, width2, width == width2)
}

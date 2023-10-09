// @Author: Ciusyan 10/8/23

package day_19

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/description/

const (
	// NIL nil 字符
	NIL = "#"
	// SPLIT 分隔符
	SPLIT = ","
)

type Codec struct {
	// 用于反序列化时使用的索引
	idx int
}

func Constructor() Codec {
	return Codec{}
}

// 层序遍历序列化
func (this *Codec) serialize(root *TreeNode) string {
	return this.levelSerialize(root)
}

// 使用层序遍历的方式反序列化
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// 裁剪字符串后传入
	return this.levelDeserialize(strings.Split(data, SPLIT))
}

// 使用层序遍历的方式序列化
func (this *Codec) levelSerialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	// 用于序列化字符
	sb := make([]rune, 0, 1)

	// 用于层序遍历的队列
	queue := NewQueue()
	queue.Offer(root)
	// 现将根节点序列化了
	sb = append(sb, []rune(fmt.Sprintf("%d%s", root.Val, SPLIT))...)

	for queue.Size() != 0 {
		node := queue.Poll()

		if node.Left != nil {
			// 层序遍历
			queue.Offer(node.Left)
			// 序列化子节点
			sb = append(sb, []rune(fmt.Sprintf("%d%s", node.Left.Val, SPLIT))...)
		} else {
			sb = append(sb, []rune(NIL+SPLIT)...)
		}

		if node.Right != nil {
			queue.Offer(node.Right)
			// 序列化子节点
			sb = append(sb, []rune(fmt.Sprintf("%d%s", node.Right.Val, SPLIT))...)
		} else {
			sb = append(sb, []rune(NIL+SPLIT)...)
		}
	}

	return string(sb[0 : len(sb)-1])
}

// 使用层序遍历的方式反序列化
func (this *Codec) levelDeserialize(data []string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	// 将根节点的值弹出来
	rootVal, _ := strconv.Atoi(data[this.idx])
	this.idx++

	// 先将根节点建起来
	root := &TreeNode{Val: rootVal}
	// 准备一个队列用于层序遍历
	queue := NewQueue()
	queue.Offer(root)

	for queue.Size() != 0 {
		// 每次弹出最前面的两个字符
		leftStr := data[this.idx]
		this.idx++
		rightStr := data[this.idx]
		this.idx++

		// 弹出队头元素
		node := queue.Poll()

		// 看左子树
		if leftStr != NIL {
			leftVal, _ := strconv.Atoi(leftStr)
			node.Left = &TreeNode{Val: leftVal}
			// 层序遍历，有左加左
			queue.Offer(node.Left)
		}

		// 看右子树
		if rightStr != NIL {
			rightVal, _ := strconv.Atoi(rightStr)
			node.Right = &TreeNode{Val: rightVal}
			// 层序遍历，有右加右
			queue.Offer(node.Right)
		}
	}

	return root
}

// Serializes a tree to a single string. 前序遍历
func (this *Codec) serialize1(root *TreeNode) string {
	sb := make([]rune, 0)
	this.preorderSerialize(root, &sb)

	return string(sb)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize1(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// 裁剪字符串
	strs := strings.Split(data, SPLIT)
	if len(strs) <= 1 {
		return nil
	}
	// strs 需要去掉最后一个 ,
	return this.preorderDeserialize(strs[0 : len(strs)-1])
}

// 按照前序遍历的方式序列化
// root 根节点，sb 用于拼接的字符串
func (this *Codec) preorderSerialize(root *TreeNode, sb *[]rune) {
	if root == nil {
		// 说明需要拼接 nil 节点
		*sb = append(*sb, []rune(NIL+SPLIT)...)
		return
	}

	// 需要先拼接值
	*sb = append(*sb, []rune(fmt.Sprintf("%d%s", root.Val, SPLIT))...)
	// 然后对左右子树也进行前序遍历
	this.preorderSerialize(root.Left, sb)
	this.preorderSerialize(root.Right, sb)
}

// 前序遍历的方式反序列化
func (this *Codec) preorderDeserialize(data []string) *TreeNode {
	if data[this.idx] == NIL {
		return nil
	}

	// 构建根节点
	val, _ := strconv.Atoi(data[this.idx])
	root := &TreeNode{Val: val}
	// 构建好一个节点索引就 +1
	this.idx++
	root.Left = this.preorderDeserialize(data)
	// 构建好一个节点索引就 +1
	this.idx++
	root.Right = this.preorderDeserialize(data)

	return root
}

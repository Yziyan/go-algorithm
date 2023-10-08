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
	// 分隔符
	SPLIT = ","
)

type Codec struct {
	// 用于反序列化时使用的索引
	idx int
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := make([]rune, 0)
	this.preorderSerialize(root, &sb)

	return string(sb)
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

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
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

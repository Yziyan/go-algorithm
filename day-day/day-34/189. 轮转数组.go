// @Author: Ciusyan 3/12/24

package day_34

// https://leetcode.cn/problems/rotate-array/

func rotate(nums []int, k int) {
	if nums == nil || len(nums) < 2 || k < 1 {
		return
	}

	l := len(nums)
	k %= l

	// 优先对前后两节进行逆序，比如 [1,2,3, 5,9,4,1] k = 3, l = 7
	reverse(nums, 0, l-k) // eg: [5,3,2,1, 9,4,1]
	reverse(nums, l-k, l) // eg: [5,3,2,1, 1,4,9]
	// 再对整体逆序
	reverse(nums, 0, l) // eg: [9,4,1, 1,2,3,5]
}

// 对 nums 的 [begin, end) 区间进行逆序
func reverse(nums []int, begin, end int) {
	// 不包含 end 位置
	end--
	for begin < end {
		nums[begin], nums[end] = nums[end], nums[begin]
		begin++
		end--
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		cur = head
		pre *ListNode
	)

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

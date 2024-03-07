// @Author: Ciusyan 3/5/24

package day_33

// https://leetcode.cn/problems/sort-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	var (
		pre      *ListNode // 每次 merge 一组后，需要连在前一组的末尾
		h        = head    // 最终的头节点
		teamHead = head    // 分组的开始节点
	)

	// 步长从 1 开始，每次增长一倍
	for step := 1; step < n; step <<= 1 {
		// 依次跑完所有组
		for teamHead != nil {
			lh, lt, rh, rt, nextTeam := hthtn(teamHead, step)
			mh, mt := merge(lh, lt, rh, rt)

			if h == teamHead {
				// 说明是第一组
				h = mh
			} else {
				// 说明是后面的组
				pre.Next = mh // 当前组合并后的头，接上上一组的尾部
			}
			// 去下一组前，将当前组合并后的尾部，给到 pre，稍后接下一组的头
			pre = mt
			// 去下一组
			teamHead = nextTeam
		}
		// 跑完一次完整的步长后，回到开头位置，去跑下一步长
		teamHead = h
		pre = nil
	}

	return h
}

// 传入从哪个头节点开始，步长为多少，得到:左头左尾 右头右尾 下一组的起始位置
func hthtn(head *ListNode, step int) (lh, lt, rh, rt, next *ListNode) {
	lh = head
	lt = head

	pass := 0 // 记录串过了几个节点
	for head != nil {
		pass++
		if pass <= step {
			// 说明左边都还没有穿过
			lt = head
		}
		if pass == step+1 {
			// 说明刚好到右边的第一个位置
			rh = head
		}
		if pass > step {
			// 说明左边找完了，右边每遇到一个，都可能是右边的终止
			rt = head
		}

		if pass == step<<1 {
			// 说明左右两边都找齐了 step 个
			break
		}

		head = head.Next
	}
	// 把左右两组的关系先断干净
	lt.Next = nil
	if rt != nil {
		// 说明右组也是满的，说明可能有下一组
		next = rt.Next
		// 右边也得断干净
		rt.Next = nil
	}

	return
}

// 传入：左头左尾 右头右尾，得到整体的头和尾
func merge(lh, lt, rh, rt *ListNode) (head, tail *ListNode) {
	if rh == nil {
		// 说明没有右边组，直接返回左边
		return lh, lt
	}

	// 说明需要合并
	var (
		pre *ListNode // cur 的前一个该接谁
		cur *ListNode // 当前应该接入的节点
	)

	// 只要左右两边都有值，就去归并比较
	for lh != lt.Next && rh != rt.Next {
		if lh.Val <= rh.Val {
			// 说明左边的小
			cur = lh
			lh = lh.Next
		} else {
			// 说明右边的小
			cur = rh
			rh = rh.Next
		}

		// 然后看看应该将 cur 接在哪里
		if pre == nil {
			// 说明这一组刚开始
			head = cur
			pre = cur
		} else {
			// 说明之前接过了
			pre.Next = cur
			pre = cur
		}
	}

	// 来到这里，至少一边肯定接完了
	if lh != lt.Next {
		// 说明左边没接完
		for lh != lt.Next {
			pre.Next = lh
			pre = lh
			tail = lh
			lh = lh.Next
		}
	} else {
		// 说明右边没接完
		for rh != rt.Next {
			pre.Next = rh
			pre = rh
			tail = rh
			rh = rh.Next
		}
	}

	return
}

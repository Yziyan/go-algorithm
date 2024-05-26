// @Author: Ciusyan 5/22/24

package cycle_13_5_22_5_26

// https://leetcode.cn/problems/sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
思路重复，要对链表排序，可否借用对数组排序的算法呢？比如快排？堆排？归并排序？
对于堆排序，可能不太好做，因为没法将节点弄成索引。
那么这里使用归并来做，并且我们使用非递归实现的归并排序。

我们将链表所有元素，按照一定步长，一步一步的排序，步长分别是：[1, 2, 4, 8, 16, ..., 2^n-1]
比如链表是：9 -> 8 -> 10 -> 5 -> 8 -> 12 -> 4 -> 6 那么
step = 1：[9 -> 8] -> [10 -> 5] -> [8 -> 12] -> [4 -> 6] 结果是：[8 -> 9] -> [5 -> 10] -> [8 -> 12] -> [4 -> 6]
step = 2：[8 -> 9 -> 5 -> 10] -> [8 -> 12 -> 4 -> 6] 结果是：[5 -> 8 -> 9 -> 10] -> [4 -> 6 -> 8 -> 12]
step = 4：[5 -> 8 -> 9 -> 10 -> 4 -> 6 -> 8 -> 12] 结果是：[4 -> 5 -> 6 -> 8 -> 8 -> 9 -> 10 -> 12]
拆分后，每一组有左右两条链表，每条链表最多不超过 step 个节点。然后将这两条链表合并成一条链表。
*/

func sortList(head *ListNode) *ListNode {
	// 根据步长，获取左右两组，分别返回：左头 左尾、右头 右尾、下一组的头
	hthtn := func(head *ListNode, step int) (lH, lT, rH, rT, next *ListNode) {
		lH = head
		lT = head

		pass := 0 // 用于记录走了多少步
		for head != nil {
			pass++
			if pass <= step {
				// 说明左边还没有收集完毕
				lT = head
			} else if pass == step+1 {
				// 说明前面收集完毕了，该右边了
				rH = head
			}

			if pass > step {
				// 右边没遇到的一个，都可能是右边的终止位置
				rT = head
			}

			if pass == step<<1 {
				// 说明左右两组都收集完毕了
				break
			}

			head = head.Next
		}
		// 把左右两组的关系断掉
		lT.Next = nil

		if rT != nil {
			// 说明右组也是满的，可能有下一组
			next = rT.Next
			rT.Next = nil
		}

		return lH, lT, rH, rT, next
	}

	// 合并左右两条链表，返回合并后的头和尾
	merge := func(lH, lT, rH, rT *ListNode) (mH, mT *ListNode) {
		if rH == nil {
			// 说明没有右组，不需要合并
			return lH, lT
		}
		// 需要合并，
		var (
			pre *ListNode
			cur *ListNode
		)

		// 只要左右两边都有值，就去比较看看放谁
		for lH != nil && rH != nil {
			if lH.Val <= rH.Val {
				// 左边的小
				cur = lH
				lH = lH.Next
			} else {
				// 右边小
				cur = rH
				rH = rH.Next
			}

			// 然后看看应该将 cur 接在哪里
			if pre == nil {
				// 说明第一组为 nil
				mH = cur
			} else {
				pre.Next = cur
			}
			pre = cur
		}

		// 来到这里，至少有一边被接完了
		// 说明左边没接完，接到后面
		for lH != nil {
			pre.Next = lH
			pre = lH
			mT = lH
			lH = lH.Next
		}

		for rH != nil {
			pre.Next = rH
			pre = rH
			mT = rH
			rH = rH.Next
		}

		return mH, mT
	}

	// 1. 先算出链表长度
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	var (
		resH    = head    // 最终返回的头结点
		groupH  = head    // 每一组的头节点
		preTail *ListNode // 上一组的结尾
	)

	// 分组的步长依次为 [1, 2, 4, 8, ... 2^n-1]
	for step := 1; step < n; step <<= 1 {
		// 跑完所有分组
		for groupH != nil {
			lH, lT, rH, rT, next := hthtn(groupH, step)
			mH, mT := merge(lH, lT, rH, rT)

			// 和上一组做关联
			if preTail == nil {
				// 说明这是第一组
				resH = mH
			} else {
				// 说明是后面的组，接上前面的组
				preTail.Next = mH
			}
			preTail = mT
			groupH = next
		}
		// 跑完这一组步长后，去下一组步长的时候，记得还原状态
		groupH = resH
		preTail = nil
	}

	return resH
}

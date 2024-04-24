// @Author: Ciusyan 4/24/24

package cycle_7_4_22_4_26

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
思路重复：
对于要翻转倒数第 n 个节点，核心就是要找到 倒数 第 n+1 个节点。
那么我们如果将链表的长度统计出来后，正着数 size-n-1 个节点，其实就是要找的节点了，
只不过在这个过程中，需要对一些边界做额外的控制。否则可能出现空指针的问题。
但是我们这样就至少需要扫描 n + size - n 次。能否有更好的做法呢？
这里就用到了快慢指针的解法：
1.准备快慢指针
2.让 fast 指针默认领先 slow n+1 个位置
3.fast 和 slow 同步往下走，频率也都是每次走一步。
4.只要 fast 走到结尾了，slow 肯定就是倒数第 n+1 个节点。
6.更新对应 Next 指针的引用
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var (
		slow = head
		fast = head
	)

	// 找到目标节点的前一个，先让 fast 和 slow 差距 n+1
	for i := 0; i < n+1; i++ {
		if fast == nil {
			// 说明在走的过程中，fast 为 nil 了
			return head.Next
		}
		fast = fast.Next
	}

	// 现在再让 fast 和 slow 同步走
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 当 fast 为 nil 后，其实 slow 就走到了目标节点的前一个节点
	slow.Next = slow.Next.Next

	return head
}

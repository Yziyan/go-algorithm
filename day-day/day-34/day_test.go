// @Author: Ciusyan 3/8/24

package day_34

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMissingRanges(t *testing.T) {

	ranges := findMissingRanges([]int{0, 2, 3, 10, 50, 78}, -3, 99)

	t.Log(ranges)
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(nums)
	rotate(nums, 3)
	t.Log(nums)
}

func TestFindLongestOrderedSubstring(t *testing.T) {

	testCases := []struct {
		name string
		str  string
		want string
	}{
		{name: "case1", str: "312351", want: "123"},
		{name: "case1", str: "4312355671", want: "123"},
		{name: "case1", str: "3136789234561", want: "23456"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findLongestOrderedSubstring(tc.str)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestReverseList(t *testing.T) {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	head = reverseList(head)
	t.Log(head)
}

func TestValidNumber(t *testing.T) {

	number := validNumber("1  ")
	number = validNumber(".-4")

	t.Log(number)
}

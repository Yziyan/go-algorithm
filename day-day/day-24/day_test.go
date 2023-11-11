// @Author: Ciusyan 11/9/23

package day_24

import (
	"fmt"
	"testing"
)

func TestHanoi(t *testing.T) {

	Hanoi(3)
	fmt.Println()
	Hanoi1(3)
}

func TestAllSubsquences(t *testing.T) {

	got := AllSubsquences("ABCA")
	got2 := AllSubsquencesNoRepeat("ABC")
	got3 := PrintAllPermutations("ABC")
	got4 := PrintAllPermutations1("ABC")
	got5 := PrintAllPermutationsNoRepeat("ABC")
	got6 := PrintAllPermutationsNoRepeat("AAC")
	t.Log(got, len(got))
	t.Log(got2, len(got2))
	t.Log(got3, len(got3))
	t.Log(got4, len(got4))
	t.Log(got5, len(got5))
	t.Log(got6, len(got6))
}

func TestReverseStackUsingRecursive(t *testing.T) {

	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	t.Log(stack)
	ReverseStackUsingRecursive(stack)
	t.Log(stack)
}

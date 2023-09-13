// @Author: Ciusyan 9/13/23

package day_8

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := &queue{}

	s.Offer(1)
	s.Offer(2)
	s.Offer(3)

	fmt.Println(s.Poll())
	fmt.Println(s.Poll())
	fmt.Println(s.Poll())
}

func TestMyStack(t *testing.T) {
	s := Constructor()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}

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
	t.Log(got, len(got))
	t.Log(got2, len(got2))
}

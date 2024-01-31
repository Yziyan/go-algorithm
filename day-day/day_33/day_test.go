// @Author: Ciusyan 1/31/24

package day_33

import "testing"

func TestIsPalindrome(t *testing.T) {
	isPalindrome("A man, a plan, a canal: Panama")
	isPalindrome(".,")
	isPalindrome("0P")
}

func TestLadderLength(t *testing.T) {

	length := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	length = ladderLength("ymain", "oecij", []string{"ymann", "yycrj", "oecij",
		"ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", "ymain"})

	t.Log(length)
}

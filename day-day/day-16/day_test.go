// @Author: Ciusyan 9/25/23

package day_16

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("ciusyan")
	trie.Insert("ciuszhiyan")
	trie.Insert("zhiyan")
	trie.Insert("zhirong")
	if !trie.Search("ciusyan") {
		t.Error("出错了")
	}
	if trie.Search("john") {
		t.Error("出错了")
	}
	trie.Delete("ciusyan")
	if trie.Search("ciusyan") {
		t.Error("出错了")
	}
	if !trie.StartsWith("cius") {
		t.Error("出错了")
	}
	trie.Delete("ciuszhiyan")
	if trie.StartsWith("cius") {
		t.Error("出错了")
	}
	if !trie.StartsWith("zhi") {
		t.Error("出错了")
	}
	if trie.StartsWith("deqiang") {
		t.Error("出错了")
	}
}

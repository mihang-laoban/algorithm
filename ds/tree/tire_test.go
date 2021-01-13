package tree

import (
	"fmt"
	"testing"
)

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		if node.next[v-'a'] == nil {
			node.next[v-'a'] = &Trie{}
		}
		node = node.next[v-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		node = node.next[v-'a']
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		node = node.next[v-'a']
		if node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);

来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func TestTire(t *testing.T) {
	tire := Constructor()
	tire.Insert("apple")
	fmt.Println(tire.Search("apple"))
	fmt.Println(tire.Search("app"))
	fmt.Println(tire.StartsWith("app"))
	tire.Insert("app")
	fmt.Println(tire.Search("app"))
}

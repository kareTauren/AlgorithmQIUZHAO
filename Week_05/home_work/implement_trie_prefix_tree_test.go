package home_work

import (
	"fmt"
	"testing"
)

type Trie struct {
	isEnd bool
	Next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.Next[v-'a'] == nil {
			this.Next[v-'a'] = new(Trie)
		}
		this = this.Next[v-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.Next[v-'a'] == nil {
			return false
		}
		this = this.Next[v-'a']
	}
	return this.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.Next[v-'a'] == nil {
			return false
		}
		this = this.Next[v-'a']
	}
	return true
}

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("any")
	trie.Insert("anyone")
	trie.Insert("anybody")
	fmt.Println(trie.Search("apple"))   // 返回 true
	fmt.Println(trie.Search("app"))     // 返回 false
	fmt.Println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	trie.Search("app")
}

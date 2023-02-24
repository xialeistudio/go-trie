// Package trie
// Author: xialeistudio<xialeistudio@gmail.com>
// Date: 2023/2/24
package trie

type Trie struct {
	Next  map[rune]*Trie
	IsEnd bool
}

// New Construct a new Trie
func New() *Trie {
	return &Trie{
		Next:  make(map[rune]*Trie),
		IsEnd: false,
	}
}

// Add word to trie
func (t *Trie) Add(word string) {
	node := t
	runes := []rune(word)
	for _, r := range runes {
		if _, ok := node.Next[r]; !ok {
			node.Next[r] = New()
		}
		node = node.Next[r]
	}
	node.IsEnd = true
}

func (t *Trie) Exists(word string) bool {
	node := t
	runes := []rune(word)
	for _, r := range runes {
		if _, ok := node.Next[r]; !ok {
			return false
		}
		node = node.Next[r]
	}
	return node.IsEnd
}

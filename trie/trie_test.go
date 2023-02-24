// Package trie
// Author: stonexia<stonexia@tencent.com>
// Date: 2023/2/24
package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Add(t1 *testing.T) {
	a := assert.New(t1)

	trie := New()
	trie.Add("敏感词")
	trie.Add("abc")
	a.True(trie.Exists("敏感词"))
	a.False(trie.Exists("敏感"))
	a.True(trie.Exists("abc"))
	a.False(trie.Exists("abcd"))
}

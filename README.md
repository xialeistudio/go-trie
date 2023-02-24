# go-trie

A Trie implementation which provides a very fast speed to check for specific word.

Supported unicode words.


## QuickStart

1. Fetch latest package

```shell
go get github.com/xialeistudio/go-trie
```
2. Using
```go
trie:= trie.New()
// add words
trie.Add("sensitive")
trie.Add("敏感词")
// search words
fmt.Println(trie.Exists("sensitive")) => true
fmt.Println(trie.Exists("敏感词")) => true
fmt.Println(trie.Exists("敏感")) => false

```
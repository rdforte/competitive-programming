package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("A -----------------")
	a := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"}, []int{5, 3, 2, 2})
	fmt.Println(a.Input('i'))
	fmt.Println(a.Input(' '))
	fmt.Println(a.Input('a'))
	fmt.Println(a.Input('#'))

	fmt.Println("B -----------------")
	b := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"}, []int{5, 3, 2, 2})
	fmt.Println(b.Input('i'))
	fmt.Println(b.Input(' '))
	fmt.Println(b.Input('a'))
	fmt.Println(b.Input('#'))
	fmt.Println(b.Input('i'))
	fmt.Println(b.Input(' '))
	fmt.Println(b.Input('a'))
	fmt.Println(b.Input('#'))
	fmt.Println(b.Input('i'))
	fmt.Println(b.Input(' '))
	fmt.Println(b.Input('a'))
	fmt.Println(b.Input('#'))
}

type AutocompleteSystem struct {
	root        *TrieNode
	curNode     *TrieNode
	curSentence []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	root := &TrieNode{
		children: make(map[byte]*TrieNode),
	}
	for i, s := range sentences {
		root.AddTrieNode(s, times[i])
	}
	return AutocompleteSystem{root, root, nil}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.curNode = this.root
		this.root.AddTrieNode(string(this.curSentence), 1)
		this.curSentence = nil
		return nil
	}

	this.curSentence = append(this.curSentence, c)

	if this.curNode == nil {
		return nil
	}

	this.curNode = this.curNode.children[c]

	if this.curNode == nil {
		return nil
	}

	// update my Trie with cur sentence

	// create heap
	h := &SentenceHeap{}
	heap.Init(h)

	// add sentences from trie node to heap
	for sent, count := range this.curNode.sentences {
		heap.Push(h, Sentence{sent, count})
	}

	res := make([]string, 0, 3)
	for i := 0; i < 3 && h.Len() > 0; i++ {
		s := heap.Pop(h).(Sentence)
		res = append(res, s.sentence)
	}

	return res
}

type TrieNode struct {
	children  map[byte]*TrieNode
	sentences map[string]int
}

func (t *TrieNode) AddTrieNode(sentence string, count int) {
	node := t
	for i := 0; i < len(sentence); i++ {
		char := sentence[i]
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{
				children:  make(map[byte]*TrieNode),
				sentences: make(map[string]int),
			}
		}
		node = node.children[char]
		node.sentences[sentence] += count
	}
}

type SentenceHeap []Sentence

func (h SentenceHeap) Len() int { return len(h) }

func (h SentenceHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		if comp := strings.Compare(h[i].sentence, h[j].sentence); comp == 1 {
			return false
		} else if comp == -1 {
			return true
		}
	}
	return h[i].count > h[j].count
}

func (h SentenceHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *SentenceHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Sentence))
}

func (h *SentenceHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Sentence struct {
	sentence string
	count    int
}

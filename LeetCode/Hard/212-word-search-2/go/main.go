package main

import "fmt"

func main() {
	fmt.Println(findWordsV2([][]byte{
		{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain"}))

	fmt.Println(findWordsV2([][]byte{
		{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'},
	}, []string{"oa", "oaa"}))
}

func findWordsV2(board [][]byte, words []string) []string {
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	rowLen := len(board)
	colLen := len(board[0])

	var res []string
	trie := NewTrie(words)

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			var dfs func(row, col int, n *TrieNode) bool
			dfs = func(row, col int, n *TrieNode) bool {
				if row >= rowLen || row < 0 || col >= colLen || col < 0 || board[row][col] == '*' {
					return false
				}

				if _, ok := n.children[board[row][col]]; !ok {
					return false
				}

				if len(n.children[board[row][col]].words) > 0 {
					res = append(res, n.children[board[row][col]].words...)
					n.children[board[row][col]].words = nil
				}

				char := board[row][col]
				board[row][col] = '*'

				for _, dir := range dirs {
					if dfs(row+dir[0], col+dir[1], n.children[char]) {
						board[row][col] = char
						return true
					}
				}

				board[row][col] = char

				return false
			}

			dfs(r, c, trie)
		}
	}

	return res
}

func NewTrie(words []string) *TrieNode {
	root := &TrieNode{
		children: make(map[byte]*TrieNode),
	}

	curNode := root
	for _, word := range words {
		for i := range word {
			char := word[i]
			if _, ok := curNode.children[char]; !ok {
				curNode.children[char] = &TrieNode{children: make(map[byte]*TrieNode)}
			}
			curNode = curNode.children[char]
		}
		curNode.words = append(curNode.words, word)
		curNode = root
	}

	return root
}

type TrieNode struct {
	children map[byte]*TrieNode
	words    []string
}

// Brute Force
// Time Limit Exceeded -----------------------------------------------------------------------------------------------------------
func findWords(board [][]byte, words []string) []string {
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	alreadyFound := make([]bool, len(words))

	rowLen := len(board)
	colLen := len(board[0])

	var res []string

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			var findWord func(row, col, i, w int) bool
			findWord = func(row, col, i, w int) bool {
				if row >= rowLen || row < 0 || col >= colLen || col < 0 || board[row][col] == '*' {
					return false
				}

				if board[row][col] != words[w][i] {
					return false
				}

				if i == len(words[w])-1 {
					return true
				}

				char := board[row][col]
				board[row][col] = '*'

				for _, dir := range dirs {
					if findWord(row+dir[0], col+dir[1], i+1, w) {
						board[row][col] = char
						return true
					}
				}

				board[row][col] = char

				return false
			}
			for i, word := range words {
				if !alreadyFound[i] && findWord(r, c, 0, i) {
					res = append(res, word)
					alreadyFound[i] = true
				}
			}
		}
	}

	return res
}

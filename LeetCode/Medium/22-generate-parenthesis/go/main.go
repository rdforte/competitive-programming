package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res := []string{}

	var genParenDFS func(p []byte, l, r int)
	genParenDFS = func(p []byte, l, r int) {
		if r > l || r > n || l > n {
			return
		}

		if len(p) == 2*n {
			res = append(res, string(p))
			return
		}

		p = append(p, '(')
		genParenDFS(p, l+1, r)
		p = p[:len(p)-1]
		p = append(p, ')')
		genParenDFS(p, l, r+1)
		p = p[:len(p)-1]
	}

	genParenDFS([]byte{}, 0, 0)

	return res
}

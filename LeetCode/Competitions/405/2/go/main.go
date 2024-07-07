package main

import "fmt"

func main() {
	fmt.Println(validStrings(3))
	fmt.Println(validStrings(5))
	fmt.Println(validStrings(1))
}

func validStrings(n int) []string {
	var res []string
	var dfs func(b []byte)
	dfs = func(b []byte) {
		if len(b) == n {
			res = append(res, string(b))
			return
		}

		if len(b) == 0 || b[len(b)-1] != '0' {
			b = append(b, '0')
			dfs(b)
			b[len(b)-1] = '1'
			dfs(b)
			b = b[:len(b)-1]
		} else {
			b = append(b, '1')
			dfs(b)
			b = b[:len(b)-1]
		}
	}

	dfs([]byte{})

	return res
}

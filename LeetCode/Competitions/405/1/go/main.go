package main

import "fmt"

func main() {
	fmt.Println(getEncryptedString("dart", 3))
	fmt.Println(getEncryptedString("darts", 3))
	fmt.Println(getEncryptedString("aaa", 1))
	fmt.Println(getEncryptedString("ab", 1))
	fmt.Println(getEncryptedString("byd", 4))
}

func getEncryptedString(s string, k int) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		index := i + k
		if index >= len(s) {
			res[i] = s[((i+k)-len(s))%len(s)]
		} else {
			res[i] = s[i+k]
		}
	}

	return string(res)
}

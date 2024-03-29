package main

import "fmt"

func main() {
	c := Codec{}
	enc := c.Encode([]string{"hello", "world"})
	fmt.Println(enc)
	strs := c.Decode(enc)

	fmt.Println(strs)
	fmt.Println(len(strs))
}

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	s := []byte{}
	for _, str := range strs {
		s = append(s, 2)
		s = append(s, str...)
		s = append(s, 3)
	}
	return string(s)
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	decodedStrings := []string{}
	byteStr := []rune{}

	for _, c := range strs {
		if c == 2 {
			byteStr = []rune{}
		} else if c == 3 {
			decodedStrings = append(decodedStrings, string(byteStr))
		} else {
			byteStr = append(byteStr, c)
		}
	}

	return decodedStrings
}


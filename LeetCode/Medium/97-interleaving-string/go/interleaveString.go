package interleaveString

import "fmt"

func IsInterleave(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	inter := Interleave{mem: map[string]bool{}, s1: s1, s2: s2, s3: s3}
	return inter.recursiveIsInterleave(0, 0, 0)
}

type Interleave struct {
	mem map[string]bool
	s1  string
	s2  string
	s3  string
}

func (i *Interleave) recursiveIsInterleave(s1Index, s2Index, s3Index int) bool {
	if s1Index == len(i.s1) && s2Index == len(i.s2) {
		return true
	}

	if val, ok := i.mem[fmt.Sprintf("%d %d", s1Index, s2Index)]; ok {
		return val
	}

	var isInter = false
	if s1Index < len(i.s1) && !isInter && i.s3[s3Index] == i.s1[s1Index] {
		isInter = i.recursiveIsInterleave(s1Index+1, s2Index, s3Index+1)
	}

	if s2Index < len(i.s2) && !isInter && i.s3[s3Index] == i.s2[s2Index] {
		isInter = i.recursiveIsInterleave(s1Index, s2Index+1, s3Index+1)
	}

	i.mem[fmt.Sprintf("%d %d", s1Index, s2Index)] = isInter

	return isInter
}

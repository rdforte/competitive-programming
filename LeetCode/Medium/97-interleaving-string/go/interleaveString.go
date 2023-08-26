package interleaveString

func IsInterleave(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	inter := Interleave{mem: map[string]string{}, s1: s1, s2: s2, s3: s3}
	return inter.recursiveIsInterleave("", 0, 0)
}

type Interleave struct {
	mem map[string]string
	s1  string
	s2  string
	s3  string
}

func (i *Interleave) recursiveIsInterleave(interleave string, s1Index, s2Index int) bool {
	if len(interleave) == len(i.s3) {
		return interleave == i.s3
	}

	var isInter = false
	if s1Index < len(i.s1) && !isInter {
		isInter = i.recursiveIsInterleave(interleave+string(i.s1[s1Index]), s1Index+1, s2Index)
	}

	if s2Index < len(i.s2) && !isInter {
		isInter = i.recursiveIsInterleave(interleave+string(i.s2[s2Index]), s1Index, s2Index+1)
	}

	return isInter
}

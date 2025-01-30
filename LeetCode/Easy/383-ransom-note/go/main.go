package main

func main() {
}

func canConstruct(ransomNote, magazine string) bool {
	magazineChars := make([]int, 26)
	for i := range magazine {
		c := magazine[i]
		magazineChars[c-'a']++
	}

	ransomNoteChars := make([]int, 26)
	for i := range ransomNote {
		c := ransomNote[i]
		ransomNoteChars[c-'a']++
	}

	for i, count := range ransomNoteChars {
		if count > magazineChars[i] {
			return false
		}
	}

	return true
}

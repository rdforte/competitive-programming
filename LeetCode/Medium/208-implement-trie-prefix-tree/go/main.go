package main

import "fmt"

func main() {
	t := Constructor()
	t.Insert("cat")
	fmt.Println(t.Search("cat"))
	fmt.Println(t.Search("ca"))
	fmt.Println(t.Search("cats"))
	t.Insert("cats")
	fmt.Println(t.Search("cats"))
	fmt.Println(t.StartsWith("ca"))
	fmt.Println(t.StartsWith("ba"))
}

func Constructor() Trie {
	characters := make([]*Character, 26)
	character := &Character{characters, false}
	return Trie{character}
}

type Trie struct {
	character *Character
}

func (this *Trie) Insert(word string) {

	var insertRecursive func(character *Character, i int)
	insertRecursive = func(character *Character, i int) {

		if character.characters[word[i]-'a'] == nil {
			characters := make([]*Character, 26)
			character.characters[word[i]-'a'] = &Character{characters, false}
		}

		if i == len(word)-1 {
			character.characters[word[i]-'a'].isWord = true
			return
		}

		insertRecursive(character.characters[word[i]-'a'], i+1)
	}

	insertRecursive(this.character, 0)
}

func (this *Trie) Search(word string) bool {

	var findWordRecursive func(character *Character, i int) bool
	findWordRecursive = func(character *Character, i int) bool {
		if character.characters[word[i]-'a'] == nil {
			return false
		}

		if i == len(word)-1 {
			return character.characters[word[i]-'a'].isWord
		}

		return findWordRecursive(character.characters[word[i]-'a'], i+1)

	}

	return findWordRecursive(this.character, 0)
}

func (this *Trie) StartsWith(prefix string) bool {

	var findPrefixRecursive func(character *Character, i int) bool
	findPrefixRecursive = func(character *Character, i int) bool {
		if character.characters[prefix[i]-'a'] == nil {
			return false
		}

		if i == len(prefix)-1 {
			return true
		}

		return findPrefixRecursive(character.characters[prefix[i]-'a'], i+1)

	}

	return findPrefixRecursive(this.character, 0)
}

type Character struct {
	characters []*Character
	isWord     bool
}

package main

import "fmt"

/*
We are working on a security system for a badged-access room in our company's building.

Given an ordered list of employees who used their badge to enter or exit the room, write a function that returns two collections:

1. All employees who didn't use their badge while exiting the room - they recorded an enter without a matching exit. (All employees are required to leave the room before the log ends.)

2. All employees who didn't use their badge while entering the room - they recorded an exit without a matching enter. (The room is empty when the log begins.)

Each collection should contain no duplicates, regardless of how many times a given employee matches the criteria for belonging to it.

ROOM IS EMPTY
records1 = {
  {"Paul",     "enter"},
MISSING ENTER for Pauline
  {"Pauline",  "exit"},
MISSING EXIT for Paul
  {"Paul",     "enter"},
  {"Paul",     "exit"},
  {"Martha",   "exit"},
  {"Joe",      "enter"},
  {"Martha",   "enter"},
  {"Steve",    "enter"},
  {"Martha",   "exit"},
  {"Jennifer", "enter"},
  {"Joe",      "enter"},
  {"Curtis",   "exit"},
  {"Curtis",   "enter"},
  {"Joe",      "exit"},
  {"Martha",   "enter"},
  {"Martha",   "exit"},
  {"Jennifer", "exit"},
  {"Joe",      "enter"},
  {"Joe",      "enter"},
  {"Martha",   "exit"},
  {"Joe",      "exit"},
MISSING ENTER for Joe
  {"Joe",      "exit"}
MISSING EXIT for Curtis
}
ROOM IS EMPTY

Expected output: {"Steve", "Curtis", "Paul", "Joe"}, {"Martha", "Pauline", "Curtis", "Joe"}

Other test cases:

records2 = {
  {"Paul", "enter"},
  {"Paul", "exit"},
}

Expected output: {}, {}

ROOM IS EMPTY
records3 = {
  {"Paul", "enter"},
missing EXIT
  {"Paul", "enter"},
  {"Paul", "exit"},
missing ENTER
  {"Paul", "exit"},
}
ROOM IS EMPTY

Expected output: {"Paul"}, {"Paul"}

records4 = {
  {"Raj", "enter"},
  {"Paul", "enter"},
  {"Paul", "exit"},
  {"Paul", "exit"},
  {"Paul", "enter"},
  {"Raj", "enter"},
}

Expected output: {"Raj", "Paul"}, {"Paul"}

All Test Cases:
mismatches(records1) => {"Steve", "Curtis", "Paul", "Joe"}, {"Martha", "Pauline", "Curtis", "Joe"}
mismatches(records2) => {}, {}
mismatches(records3) => {"Paul"}, {"Paul"}
mismatches(records4) => {"Raj", "Paul"}, {"Paul"}

n: length of the badge records array
*/

func main() {
	records1 := [][]string{
		{"Paul", "enter"},
		{"Pauline", "exit"},
		{"Paul", "enter"},
		{"Paul", "exit"},
		{"Martha", "exit"},
		{"Joe", "enter"},
		{"Martha", "enter"},
		{"Steve", "enter"},
		{"Martha", "exit"},
		{"Jennifer", "enter"},
		{"Joe", "enter"},
		{"Curtis", "exit"},
		{"Curtis", "enter"},
		{"Joe", "exit"},
		{"Martha", "enter"},
		{"Martha", "exit"},
		{"Jennifer", "exit"},
		{"Joe", "enter"},
		{"Joe", "enter"},
		{"Martha", "exit"},
		{"Joe", "exit"},
		{"Joe", "exit"},
	}
	fmt.Println(mismatches(records1))
	fmt.Println(mismatchesV2(records1))

	records2 := [][]string{
		{"Paul", "enter"},
		{"Paul", "exit"},
	}
	fmt.Println(mismatches(records2))
	fmt.Println(mismatchesV2(records2))

	records3 := [][]string{
		{"Paul", "enter"},
		{"Paul", "enter"},
		{"Paul", "exit"},
		{"Paul", "exit"},
	}
	fmt.Println(mismatches(records3))
	fmt.Println(mismatchesV2(records3))

	records4 := [][]string{
		{"Raj", "enter"},
		{"Paul", "enter"},
		{"Paul", "exit"},
		{"Paul", "exit"},
		{"Paul", "enter"},
		{"Raj", "enter"},
	}
	fmt.Println(mismatches(records4))
	fmt.Println(mismatchesV2(records4))

	records5 := [][]string{
		{"Raj", "exit"},
	}
	fmt.Println(mismatches(records5))
	fmt.Println(mismatchesV2(records5))
}

func mismatchesV2(entries [][]string) ([]string, []string) {
	nameEtries := map[string][]string{}

	for _, e := range entries {
		name := e[0]
		state := e[1]
		nameEtries[name] = append(nameEtries[name], state)
	}

	didntExit := map[string]struct{}{}
	didntEnter := map[string]struct{}{}

	for name, e := range nameEtries {
		for i := 0; i < len(e); i++ {
			if i == 0 && e[i] == "exit" {
				didntEnter[name] = struct{}{}
				continue
			}

			if e[i] == "exit" && i > 0 && e[i-1] == "exit" {
				didntEnter[name] = struct{}{}
				continue
			}

			if e[i] == "enter" && i > 0 && e[i-1] == "enter" {
				didntExit[name] = struct{}{}
				continue
			}

			if i == len(e)-1 && e[i] == "enter" {
				didntExit[name] = struct{}{}
			}
		}
	}

	enterRes := make([]string, 0, len(didntEnter))
	for key := range didntEnter {
		enterRes = append(enterRes, key)
	}

	exitRes := make([]string, 0, len(didntExit))
	for key := range didntExit {
		exitRes = append(exitRes, key)
	}

	return exitRes, enterRes
}

func mismatches(entries [][]string) ([]string, []string) {
	hasEnter := map[string]struct{}{}

	didntExit := map[string]struct{}{}
	didntEnter := map[string]struct{}{}

	for _, e := range entries {
		name := e[0]
		state := e[1]

		if state == "enter" {
			if _, ok := hasEnter[name]; ok {
				didntExit[name] = struct{}{}
			} else {
				hasEnter[name] = struct{}{}
			}
		} else {
			if _, ok := hasEnter[name]; !ok {
				didntEnter[name] = struct{}{}
			}
			delete(hasEnter, name)
		}
	}

	for name := range hasEnter {
		didntExit[name] = struct{}{}
	}

	enterRes := make([]string, 0, len(didntEnter))
	for key := range didntEnter {
		enterRes = append(enterRes, key)
	}

	exitRes := make([]string, 0, len(didntExit))
	for key := range didntExit {
		exitRes = append(exitRes, key)
	}

	return exitRes, enterRes
}

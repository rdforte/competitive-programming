/*
We are working on a security system for a badged-access room in our company's building.

Given an ordered list of employees who used their badge to enter or exit the room, write a function that returns two collections:

1. All employees who didn't use their badge while exiting the room - they recorded an enter without a matching exit. (All employees are required to leave the room before the log ends.)

2. All employees who didn't use their badge while entering the room - they recorded an exit without a matching enter. (The room is empty when the log begins.)

Each collection should contain no duplicates, regardless of how many times a given employee matches the criteria for belonging to it.

ROOM IS EMPTY
records1 = [
  ["Paul",     "enter"],
MISSING ENTER for Pauline
  ["Pauline",  "exit"],
MISSING EXIT for Paul
  ["Paul",     "enter"],
  ["Paul",     "exit"],
  ["Martha",   "exit"],
  ["Joe",      "enter"],
  ["Martha",   "enter"],
  ["Steve",    "enter"],
  ["Martha",   "exit"],
  ["Jennifer", "enter"],
  ["Joe",      "enter"],
  ["Curtis",   "exit"],
  ["Curtis",   "enter"],
  ["Joe",      "exit"],
  ["Martha",   "enter"],
  ["Martha",   "exit"],
  ["Jennifer", "exit"],
  ["Joe",      "enter"],
  ["Joe",      "enter"],
  ["Martha",   "exit"],
  ["Joe",      "exit"],
MISSING ENTER for Joe
  ["Joe",      "exit"]
MISSING EXIT for Curtis
]
ROOM IS EMPTY

Expected output: ["Steve", "Curtis", "Paul", "Joe"], ["Martha", "Pauline", "Curtis", "Joe"]

Other test cases:

records2 = [
  ["Paul", "enter"],
  ["Paul", "exit"],
]

Expected output: [], []

ROOM IS EMPTY
records3 = [
  ["Paul", "enter"],
missing EXIT
  ["Paul", "enter"],
  ["Paul", "exit"],
missing ENTER
  ["Paul", "exit"],
]
ROOM IS EMPTY

Expected output: ["Paul"], ["Paul"]

records4 = [
  ["Raj", "enter"],
  ["Paul", "enter"],
  ["Paul", "exit"],
  ["Paul", "exit"],
  ["Paul", "enter"],
  ["Raj", "enter"],
]

Expected output: ["Raj", "Paul"], ["Paul"]

All Test Cases:
mismatches(records1) => ["Steve", "Curtis", "Paul", "Joe"], ["Martha", "Pauline", "Curtis", "Joe"]
mismatches(records2) => [], []
mismatches(records3) => ["Paul"], ["Paul"]
mismatches(records4) => ["Raj", "Paul"], ["Paul"]

n: length of the badge records array
*/

package main

import "fmt"

func main() {
	// sample inputs
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

	records2 := [][]string{
		{"Paul", "enter"},
		{"Paul", "exit"},
	}

	records3 := [][]string{
		{"Paul", "enter"},
		{"Paul", "enter"},
		{"Paul", "exit"},
		{"Paul", "exit"},
	}

	records4 := [][]string{
		{"Raj", "enter"},
		{"Paul", "enter"},
		{"Paul", "exit"},
		{"Paul", "exit"},
		{"Paul", "enter"},
		{"Raj", "enter"},
	}

	// enter = paul, exit = paul, raj

	fmt.Println(findMissingRecords(records1))
	fmt.Println(findMissingRecords(records2))
	fmt.Println(findMissingRecords(records3))
	fmt.Println(findMissingRecords(records4))
}

func findMissingRecords(records [][]string) []map[string]struct{} {
	enter := make(map[string]struct{})
	exit := make(map[string]struct{})

	exitLogs := make(map[string]struct{})
	enterLogs := make(map[string]struct{})

	for _, record := range records {
		name := record[0]
		state := record[1]

		_, hasEntered := enter[name]
		if state == "enter" && hasEntered {
			exitLogs[name] = struct{}{}
		}

		if state == "exit" && !hasEntered {
			enterLogs[name] = struct{}{}
		}

		if state == "exit" {
			delete(enter, name)
			delete(exit, name)
		}

		if state == "enter" {
			enter[name] = struct{}{}
		} else {
			exit[name] = struct{}{}
		}
	}

	for key := range enter {
		exitLogs[key] = struct{}{}
	}

	return []map[string]struct{}{
		enterLogs,
		exitLogs,
	}
}

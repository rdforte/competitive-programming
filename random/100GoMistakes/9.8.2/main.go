package main

import "fmt"

// Value receivers also work on pointer receivers but not the other way around.
// Any form of string formatting using fmt will use the Stringer interface to do the formatting.

func main() {
	p := Person{
		FName: "Ryan",
		LName: "Forte",
		Age:   30,
	}

	fmt.Println(p)
	description := fmt.Sprintf("%s", p)
	fmt.Println(description)
	description = fmt.Sprintf("%v", p)
	fmt.Println(description)
	err := fmt.Errorf("error: %v", p)
	panic(err)
}

type Person struct {
	FName string
	LName string
	Age   int
}

func (p Person) String() string {
	return fmt.Sprintf("person %s %s is %d years old", p.FName, p.LName, p.Age)
}

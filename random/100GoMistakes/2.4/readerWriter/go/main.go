package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Reader -------------------------------------------------------------------------------------------------------
	input := "foo"

	dataSource := strings.NewReader(input)

	dest := make([]byte, dataSource.Len())

	i, err := dataSource.Read(dest) // read from source
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d\n", i)
	fmt.Println(string(dest))

	// Writer --------------------------------------------------------------------------------------------------------
	dataSource2 := bytes.NewBuffer(make([]byte, 0))
	dataSource2.Write([]byte("hello")) // write data to the data source

	// first read

	i, err = dataSource2.Read(dest) // read data from the data source
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d\n", i)
	fmt.Println(string(dest)) // note here we did not read everything because len of dest is 3 and len of dataSource2 is 5

	// second read

	i, err = dataSource2.Read(dest) // read data from the data source - hel
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d\n", i)
	fmt.Println(string(dest)) // this will override the first 2 characters in dest - lol

	// third read

	i, err = dataSource2.Read(dest) // read data from the buffer - hel
	if err != nil {
		fmt.Println(err) // now that we have read all the data we get EOF error
	}

	// NOTE: the above can be done in a loop and we can break when we get an EOF.

}

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	b := bytes.NewReader([]byte("Hello, World!"))
	var buf bytes.Buffer
	buf.ReadFrom(b)
	fmt.Println(buf.String())

	// custom

	r := &CustomReader{s: []byte("Hello, World!")}
	var buf1 bytes.Buffer
	buf1.ReadFrom(r)
	fmt.Println(buf1.String())

	// write

	var b2 bytes.Buffer
	b2.Write([]byte("Hello, World!"))
	fmt.Println(b2.String())

	// func

	var b3 bytes.Buffer
	getData(&b3)
	fmt.Println(b3.String())

	// init
	v4 := bytes.NewBuffer([]byte("Hello, World!"))
	fmt.Println(string(v4.Bytes()))
	str, _ := v4.ReadString(',')
	fmt.Println(str)
	fmt.Println(v4.String()) // only gets the remainder now because we have read Hello, from the buffer
}

type CustomReader struct {
	s []byte
	i int
}

func (r *CustomReader) Read(b []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return
}

func getData(w io.Writer) {
	w.Write([]byte("Hello, World!"))
}


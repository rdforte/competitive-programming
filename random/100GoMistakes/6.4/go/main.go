package main

import (
	"errors"
	"fmt"
	"strings"
)

// https://codefibershq.com/blog/golang-why-nil-is-not-always-nil

func main() {
	var p *int
	if p == nil {
		fmt.Printf("p is nil: %v\n", p)
	}

	var i interface{}
	if i == nil {
		fmt.Printf("i is nil: %v\n", i)
	}

	i = p
	if i == nil {
		fmt.Printf("i is nil: %v\n", i)
	} else {
		// interface i no longer has type nil, value nil but is instead type *int and value nil
		// for an interface to be nil it must by (type = nil, value = nil)
		fmt.Printf("i is NOT nil: %v\n", i)
	}

	// var multiError *MultiError
	// fmt.Printf("multiError: %v, %v\n", multiError, reflect.TypeOf(multiError))
	// if multiError != nil {
	// 	log.Fatalf("multiError is not nil: %v", multiError) // doesn't get executed
	// }

	// var mErr error
	// fmt.Printf("mErr: %v, %v\n", mErr, reflect.TypeOf(mErr))
	// mErr = multiError
	// if mErr != nil {
	// 	log.Fatalf("mErr is not nil: %v, %v", mErr, reflect.TypeOf(mErr)) // gets executed
	// }

	// customer := Customer{33, "John"}
	// if err := customer.Validate(); err != nil {
	// 	log.Fatalf("customer is invalid: %v", err) // gets executed
	// }

	// var err error
	// if err != nil {
	// 	log.Fatalf("err is not nil: %v", err)
	// }
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("Age is negative"))
	}

	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	return m
}

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

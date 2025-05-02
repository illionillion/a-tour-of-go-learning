package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func sum(a, b any) (int, error) {
	ai, okA := a.(int)
	bi, okB := b.(int)

	if okA && okB {
		return ai + bi, nil
	}
	return 0, &MyError{time.Now(), "sum: both arguments must be int"}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	if result, err := sum(1, 3); err != nil {
		fmt.Println(result, err)
		} else {
		fmt.Println(result, err)
	}

	if result, err := sum("1", 3); err != nil {
		fmt.Println(result, err)
		} else {
		fmt.Println(result, err)
	}

}

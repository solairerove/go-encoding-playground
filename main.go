package main

import (
	"fmt"

	m "github.com/solairerove/go-encoding-playground/marshal"
	s "github.com/solairerove/go-encoding-playground/sort"
	u "github.com/solairerove/go-encoding-playground/unmarshal"
)

func main() {
	fmt.Println("Encoding examples")

	m.JSONMarshalling()
	u.JSONUnmarshalling()

	s.Sorting()
}

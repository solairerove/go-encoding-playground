package main

import (
	"fmt"

	c "github.com/solairerove/go-encoding-playground/crypt"
	m "github.com/solairerove/go-encoding-playground/marshal"
	s "github.com/solairerove/go-encoding-playground/sort"
	u "github.com/solairerove/go-encoding-playground/unmarshal"
)

func main() {
	fmt.Println("Encoding examples")

	m.JSONMarshalling()
	u.JSONUnmarshalling()

	s.Sorting()

	c.AwesomeCrypt()
}

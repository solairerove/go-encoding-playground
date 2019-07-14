package main

import (
	"fmt"

	m "github.com/solairerove/go-encoding-playground/marshal"
	u "github.com/solairerove/go-encoding-playground/unmarshal"
)

func main() {
	fmt.Println("Encoding examples")

	m.JSONMarshalling()
	u.JSONUnmarshalling()
}

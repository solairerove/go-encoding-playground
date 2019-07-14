package marshal

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

// JSONMarshalling tbd
func JSONMarshalling() {
	fmt.Println("\nJSONMarshalling")

	p1 := person{
		First: "Tyler",
		Last:  "Durden",
		Age:   32,
	}

	p2 := person{
		First: "mikita",
		Last:  "gulag",
		Age:   23,
	}

	people := []person{
		p1, p2,
	}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

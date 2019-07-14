package unmarshal

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

// JSONUnmarshalling tbd
func JSONUnmarshalling() {
	fmt.Println("\nJSONUnmarshalling")

	s := `[
					{
						"first":"Tyler",
						"last":"Durden",
						"age":32
					},
					{
						"first":"mikita",
						"last":"gulag",
						"age":23
					}
				]`

	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range people {
		fmt.Println(p)
	}
}

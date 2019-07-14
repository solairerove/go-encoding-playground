package sort

import (
	"fmt"
	"sort"
)

// Sorting tbd
func Sorting() {
	fmt.Println("\nSorting")

	xi := []int{4, 7, 10, 99, 89, 201, 43, 42}
	xs := []string{"Tyler", "Durden", "Jack", "Marla", "Bob"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

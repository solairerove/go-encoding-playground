package sort

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type sortByName []person

func (a sortByName) Len() int           { return len(a) }
func (a sortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByName) Less(i, j int) bool { return a[i].first < a[j].first }

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

	fmt.Println("\nCustom Sorting")

	p1 := person{
		first: "Jack",
		age:   33,
	}

	p2 := person{
		first: "Tyler",
		age:   33,
	}

	p3 := person{
		first: "Marla",
		age:   33,
	}

	p4 := person{
		first: "Bob",
		age:   43,
	}

	people := []person{
		p1, p2, p3, p4,
	}

	fmt.Println(people)
	sort.Sort(sortByName(people))
	fmt.Println(people)
}

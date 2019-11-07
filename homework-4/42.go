package main

import (
		"fmt"
		"sort"
)

type Phones struct{
	Name string
	Number int
}

func (p Phones) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Number)
}

type ByNumb []Phones

func (p ByNumb) Len() int           { return len(p) }
func (p ByNumb) Swap(i,j int)       { p[i], p[j] = p[j], p[i] }
func (p ByNumb) Less(i, j int) bool { return p[i].Number < p[j].Number }


func main() {
	
	phonebook := []Phones{
		{"Ivan", 89265550101},
		{"Petr", 89165550202},
		{"Michail", 89855550303},
		{"Elena", 89175550404},

	}

	fmt.Println(phonebook)

	sort.Sort(ByNumb(phonebook))

	fmt.Println(phonebook)


}
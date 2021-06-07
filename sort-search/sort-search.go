package main

import (
	"log"
	"sort"
)

type Person struct {
	Name string
	Age int
}

// Use x < j 	=> ascending order
// Use x > j 	=> descending order
func sortByName(people []Person) {
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Name < people[j].Name
	})
}

// Use x >= key		=> when slice is in ascending order
// Use x <= key 	=> when slice is in descending order
func searchByName(people []Person, name string) int {
	index := sort.Search(len(people), func(i int) bool {
		return people[i].Name >= name
	})
	return index
}

func sortByAge(people []Person) {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

func searchByAge(people []Person, age int) int {
	index := sort.Search(len(people), func(i int) bool {
		return people[i].Age >= age
	})
	return index
}

func main() {
	var people = []Person{
		{
			Name: "rishi",
			Age: 25,
		},
		{
			Name: "micheal",
			Age: 18,
		},
		{
			Name: "Bob",
			Age: 31,
		},
		{
			Name: "John",
			Age: 17,
		},
		{
			Name: "Jehnny",
			Age: 26,
		},
	}

	log.Println(people)

	sortByName(people)
	log.Println(people)

	index := searchByName(people, "rishi")
	if index < len(people) && people[index].Name == "rishi" {
		log.Println("found at index:", index)
	} else {
		log.Println("No person by name rishi")
		log.Println("But this person can be inserted into slice at this index:", index)
	}

	sortByAge(people)
	log.Println(people)

	index = searchByAge(people, 17)
	if index < len(people) && people[index].Age == 17 {
		log.Println("found at index:", index)
	} else {
		log.Println("No person by age 17")
		log.Println("But this person can be inserted into slice at this index:", index)
	}
}
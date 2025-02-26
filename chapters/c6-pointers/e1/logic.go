package c6

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakeThousandPerson(firstName, lastName string, age int) []Person {
	start := time.Now()
	var pSlice []Person

	for i := 1; i < 10_000_000; i++ {
		pSlice = append(pSlice, MakePerson(firstName, lastName, age))
	}

	fmt.Printf("MakeThousandPerson took %s\n", time.Since(start))
	return pSlice
}

func MakeThousandPersonBetter(firstName, lastName string, age int) []Person {
	start := time.Now()
	pSlice := make([]Person, 10_000_000)

	for i := 1; i < 10_000_000; i++ {
		pSlice = append(pSlice, MakePerson(firstName, lastName, age))
	}

	fmt.Printf("MakeThousandPerson took %s\n", time.Since(start))
	return pSlice
}

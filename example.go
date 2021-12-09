package main

import "fmt"

type Person struct {
	name string
	age  int
}

func add(v1 int, v2 int) int {
	return v1 + v2
}

func (person Person) updatePersonFail(name string, age int) {
	person.name = name
	person.age = age
}

func (person *Person) updatePerson(name string, age int) {
	person.name = name
	person.age = age
}

func main() {
	fmt.Println("\nStart")
	fmt.Println("Hello")

	// Function
	fmt.Println("\nFunction")
	fmt.Println("Add: ", add(4, 5))

	// Structure
	fmt.Println("\nStructure")
	// Person before change
	person := Person{"Jimmy", 26}
	fmt.Println("Before name: ", person.name)

	// Person after change
	person.updatePersonFail("Tommy", 27) // value
	fmt.Println("After name failed: ", person.name)
	(&person).updatePerson("Tommy", 27) // reference
	fmt.Println("After name: ", person.name)

	// Array
	fmt.Println("\nArray")
	var a [5]int // Empty
	a[0] = 1
	a[1] = 2
	fmt.Println("Array a: ", a)

	b := [5]int{1, 2, 3, 4, 5} // Directly
	fmt.Println("Array b: ", b)

	c := [...]string{"Hello", "world"} // ... -> size determined by number of initialized elements
	fmt.Println("Array c: ", c)

	// Anonymous function
	func() {
		fmt.Println("Anonymous function")
	}()

	// Pointer
	fmt.Println("\nPointer")
	num := 123
	numPtr := &num
	fmt.Println("Num variable: ", num)
	fmt.Println("Num pointer", numPtr)
	fmt.Println("Num pointer variable", *numPtr)

	// If
	fmt.Println("\nIf statement")
	if 100 > 200 {
		fmt.Println("False")
	} else if 200 > 100 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// Switch
	fmt.Println("\nSwitch statement")
	d := 4
	switch d {
	case 4:
		fmt.Println(d)
	default:
		fmt.Println("No d")
	}

	// For
	fmt.Println("\nFor loop")
	for i := 0; i < 5; i++ {
		fmt.Println("For ", i)
	}

	i := 0
	for i < 3 {
		fmt.Println("Similar to while loop in c")
		i++
	}

	for _, j := range [5]string{"H", "e", "l", "l", "o"} {
		fmt.Println(j) // Similar to Python enumerate()
	}
}

package main

import "fmt"

// Just property name, and the type it should be
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// 40. Embedding Structs
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// 38. Declaring Structs
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// 39. Updating Struct Values
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// 40. Embedding Structs
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jim.updateName("jimmy")
	// jim.print()
	// 43. Structs with Pointers
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()

	// 45. Pointer Shortcut
	jim.updateName("Jimbo")
	jim.print()
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// 43. Structs with Pointers
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Struct receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}

// Homework
// package main

// import "fmt"

// func main() {
// 	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	for _, item := range ints {
// 		res := item % 2

// 		var parity string

// 		if res == 0 {
// 			parity = "even"
// 		} else {
// 			parity = "odd"
// 		}

// 		fmt.Println(item, "is", parity)
// 	}
// }

// Exercise
// package main

// import "fmt"

// func main() {
// 	name := "bill"

// 	namePointer := &name

// 	fmt.Println(&namePointer)
// 	printPointer(namePointer)
// }

// func printPointer(namePointer *string) {
// 	fmt.Println(&namePointer)
// }

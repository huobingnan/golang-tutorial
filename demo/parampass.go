package main

import "fmt"

// golang parameter pass way

type Person struct {
	Name   string
	Age    int
	Gender string
}

// pass by pointer/reference
func acceptPointer(ref *int) {
	*ref = 12345
}

// value copy
func acceptValueCopy(copyValue int) {
	copyValue = 123456
}

func structValueCopy(p Person) {
	p.Name = "#Bob#"
}

func structRef(p *Person) {
	p.Name = "#Bob#"
}

func main() {
	number := 1234
	fmt.Println("number = ", number)
	acceptPointer(&number)
	fmt.Println("number = ", number)
	acceptValueCopy(number)
	fmt.Println("number = ", number)

	p := Person{Name: "Tom", Age: 14, Gender: "Male"}
	fmt.Printf("p = %v\n", p)
	structValueCopy(p)
	fmt.Printf("p = %v\n", p)
	structRef(&p)
	fmt.Printf("p = %v\n", p)
}

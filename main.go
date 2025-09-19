package main

import (
	"encoding/json"
	"fmt"
	"math"
)

//Task1

func main() {
	// Chapter 1

	//Task1
	/*	fmt.Println("Hello world!")*/

	//Task2

	/*var a int = 12
	var b float64 = 312.42
	var c string = "Hi!"

	fmt.Println(a, b, c)*/

	// Here used :=  short variant used only on func and only if first time appears
	/*
		a := 1
		b := 24.42
		c := "String type"

		fmt.Println(a, b, c)*/

	//Task3
	// a

	/*	var someNumber int
		fmt.Println("Enter number to check positive or negative")
		fmt.Scan(&someNumber)

		if someNumber > 0 {
			fmt.Printf("positive: %d\n", someNumber)
		}
		if someNumber < 0 {
			fmt.Printf("Negative: %d\n", someNumber)
		}
		if someNumber == 0 {
			fmt.Printf("Zero: %d", someNumber)
		}
	*/
	// b

	/*
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}*/

	// c

	/*	var day int
		fmt.Println("Enter day number")
		fmt.Scan(&day)

		switch day {
		case 1:
			fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thurday")
		case 5: fmt.Println("Friday")
		case 6: fmt.Println("Saturday")
		case 7: fmt.Println("Sunday")
		}
	*/

	// Task 4

	//a

	/*	var a int = 4
		var b int = 7
		fmt.Println(add(a, b))*/

	//b
	/*
		var a string = "World!"
		var b string = "Hello"
		fmt.Println(swap(a, b))

	*/

	//c

	/*var a int
	var b int
	fmt.Println("Write num 1")
	fmt.Scanln(&a)
	fmt.Println("Write num 2")
	fmt.Scanln(&b)

	fmt.Println("quotient and remainder")
	fmt.Println(rem(a, b))
	*/

	// Chapter 2 OOP

	//Task1

	/*
		var p = Person{Name: "Rasul", Age: 21}
		Greet(p)

	*/

	//Task2

	/*var e = Employee{Name: "Rasul", ID: 0}
	var m = Manager{
		Employee:   e,
		Department: "Test",
	}
	work(m.Employee)
	*/

	//Task3
	/*
		var c Shape = Circle{Radius: 4}
		fmt.Println(c.Area())

		var r Shape = Rectangle{
			Width:  10,
			Height: 5,
		}
		fmt.Println(r.Area())
	*/

	var p Product = Product{
		name:     "Test",
		price:    100,
		quantity: 1,
	}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	_ := json.Unmarshal(data)
	fmt.Println(data)
}

type Product struct {
	name     string
	price    int
	quantity int
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	ID   int
}
type Manager struct {
	Employee
	Department string
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Greet(p Person) {
	fmt.Printf("Hello %s !  ", p.Name)
}

func work(e Employee) {
	fmt.Println(e.Name, e.ID)
}

func add(a int, b int) int {
	return a + b
}

func swap(s1 string, s2 string) string {
	return s2 + " " + s1
}

func rem(a int, b int) (int, int) {
	return a / b, a % b
}

func Greet2(p Person) {
	fmt.Printf("Hello %s !", p.Name)
}

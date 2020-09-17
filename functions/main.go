package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	sayMessage("Hello, world!")

	sayGreeting("Hello", "John")

	sum := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum", sum)

	fmt.Println("Sum1", sum1(1, 2))

	d, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

	f := func() {
		fmt.Println("Hello, Go!")
	}
	f()

	g := greeter{
		greeting: "Hello",
		name:     "Jane",
	}
	g.greet()
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result = result + v
	}
	return result
}

func sum1(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result = result + v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

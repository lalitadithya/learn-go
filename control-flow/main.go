package main

import "fmt"

func main() {
	if false {
		fmt.Println("Hello, world!")
	}

	myMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	if num, ok := myMap["a"]; ok {
		fmt.Println("Number is", num)
	}

	number := 50
	guess := 50
	if guess < 1 || guess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too less")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("You got it")
		}
	}

	switch 2 {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Number is not 1 or 2")
	}

	switch 8 {
	case 1, 5, 10:
		fmt.Println("Number is 1, 5, 10")
	case 2, 4, 8:
		fmt.Println("Number is 2, 4, 8")
	default:
		fmt.Println("Number is not 1 or 2")
	}

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("Number is 1, 5, 10")
	case 2, 4, 8:
		fmt.Println("Number is 2, 4, 8")
	default:
		fmt.Println("Number is not 1 or 2")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("<=10")
		fallthrough // will not execute the logic condition in the next case
	case i <= 20:
		fmt.Println("<=20")
	default:
		fmt.Println("<20")
	}

	var a interface{} = 1
	switch a.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float")
	default:
		fmt.Println("not known")
	}
}

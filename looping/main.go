package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while
	j := 50
	for j < 55 {
		fmt.Println(j)
		j++
	}

	k := 100
	for {
		fmt.Println(k)
		k++
		if k == 105 {
			break
		}
	}

	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}
}

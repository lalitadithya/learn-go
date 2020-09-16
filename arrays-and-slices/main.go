package main

import "fmt"

func main() {
	grades := [3]int{97, 85, 74}
	fmt.Printf("Grades: %v\n", grades)

	grades1 := [...]int{97, 85, 74}
	fmt.Printf("Grades: %v\n", grades1)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "John"
	students[1] = "Jane"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 100
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 100
	fmt.Println(c)
	fmt.Println(d)

	e := []int{100, 200, 300}
	fmt.Println(e)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(e))
	f := e
	f[1] = -1
	fmt.Println(f)
	fmt.Println(e)

	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h := g[:]
	i := g[3:]
	j := g[:6]
	k := g[3:6]
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	l := make([]int, 3)
	fmt.Println(l)
	fmt.Printf("Length: %v\n", len(l))
	fmt.Printf("Capacity: %v\n", cap(l))

	m := make([]int, 3, 100)
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))

	n := []int{}
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))
	n = append(n, 1)
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))
	n = append(n, 2, 3, 4, 5, 6)
	fmt.Println(n)
	fmt.Printf("Length: %v\n", len(n))
	fmt.Printf("Capacity: %v\n", cap(n))

	stack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(stack)
	q := stack[1:] //remove first
	fmt.Println(q)
	w := stack[:len(stack)-1] //remove last
	fmt.Println(w)
	r := append(stack[:2], stack[3:]...)
	fmt.Println(r)
	fmt.Println(stack)
}

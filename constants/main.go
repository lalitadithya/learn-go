package main

import "fmt"

const a int = 10

const c = iota

const (
	d = iota
	e = iota
	f = iota
)

const (
	i = iota
	j
	k
)

// enum (kinda)
const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	fmt.Printf("%v, %T\n", a, a)
	const a int = 14
	fmt.Printf("%v, %T\n", a, a)

	const myConst1 = 43
	fmt.Printf("%v, %T\n", myConst1, myConst1)

	const x = 50
	var y int16 = 50
	fmt.Printf("%v, %T\n", x+y, x+y) // gets converted to 50 + y

	fmt.Printf("%v, %T\n", c, c)

	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	var specialistType int = catSpecialist
	fmt.Printf("%v, %T\n", specialistType == catSpecialist, specialistType == catSpecialist)

	var specialistType1 int
	fmt.Printf("%v, %T\n", specialistType1 == catSpecialist, specialistType1 == catSpecialist)

	fileSize := 40000000000.
	fmt.Printf("%.2f GB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b, %T\n", roles, roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
}

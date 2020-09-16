package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
	}
	fmt.Println(statePopulations)

	map1 := make(map[string]int)
	map1 = map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(map1)
	map1["c"] = 3
	fmt.Println(map1["c"])
	delete(map1, "a")
	fmt.Println(map1)
	ele, ok := map1["z"]
	fmt.Println(ele, ok)
	fmt.Println(len(map1))

	// struct is a value type
	aDoctor := Doctor{
		number:    3,
		actorName: "John Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	newDoctor := struct{ name string }{name: "Tom Baker"}
	fmt.Println(newDoctor)

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

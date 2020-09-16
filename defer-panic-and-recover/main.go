package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// defer is called in lifo
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("About to panick")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// can repanick, i.e., rethrow
			panic(err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("done panicking")
}

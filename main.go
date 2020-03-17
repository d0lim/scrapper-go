package main

import (
	"fmt"

	"github.com/imdigo/scrapper-go/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	definition, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("found", "hello", "definition:", definition)
	err2 := dictionary.Add("hello", "Greeting")
	if err2 != nil {
		fmt.Println(err2)
	}
}

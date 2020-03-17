package main

import (
	"fmt"

	"github.com/imdigo/scrapper-go/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First word")
	def, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(def)
	err = dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	def, err = dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
}

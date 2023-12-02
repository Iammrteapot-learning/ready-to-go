package main

import "fmt"

// Create a new type of deck
type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

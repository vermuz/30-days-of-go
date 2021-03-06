// https://golang.org/ref/spec#Interface_types
// https://godoc.org/math/rand
// https://godoc.org/strconv

/*
* Functions with a signature of an interface will accept any value of any type that implements the interface
* let's implement our own interface
 */
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Movie struct {
	title, year, quote string
}

type Music struct {
	song, year, lyrics string
}

// function getData can change its bahaviour depending on its data type
func (m Movie) getData() map[string]string {
	movie := map[string]string{
		"title": m.title,
		"year":  m.year,
		"quote": m.quote,
	}

	return movie
}

func (s Music) getData() map[string]string {
	randomNumber := strconv.Itoa(rand.Intn(100))

	song := map[string]string{
		"song":     s.song,
		"year":     s.year,
		"lyrics":   s.lyrics,
		"listners": randomNumber,
	}

	return song
}

/*
* interfaces can define what actions our types can execute
* an empty interface can hold values of any type which we will explore later
*
* interfaces is a collection of method signatures
* it has a name and type field similiar to structs
* basiclly we're creating a interface for a method that can be
* shared amongst data types that implements the same method
 */
type StreamingService interface {
	getData() map[string]string
}

// canTakeManyTypes can take many types and change its behaviour
func canTakeManyTypes(i StreamingService) {
	fmt.Printf("%T\n", i)

	for key, value := range i.getData() {
		fmt.Printf("%v -> %v\n", key, value)
	}
}

func main() {
	matrix := Movie{
		"Matrix",
		"1999",
		"Never send a human to do a machine's job -Agent Smith",
	}

	daftPunk := Music{
		"One More Time",
		"2001",
		"One more time we're gonna celebrate, Oh yeah all right don't stop the dancing",
	}

	canTakeManyTypes(matrix)
	canTakeManyTypes(daftPunk)
}

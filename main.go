package main

import (
	"flag"
	"fmt"
)

func main() {

	note := flag.String("note", "a", "Root note to start from")
	number := flag.Int("number", 4, "Number of notes to return if circle method is chosen")
	method := flag.String("method", "", "Method to use, like circle for circle of fifths")

	flag.Parse()

	i := 0
	rootNote := *note
	notes := returnNotes()
	fmt.Println("Root Note:", rootNote)

	if *method == "circle" {
		for i < *number {
			newNote := circleOfFifths(rootNote, notes)
			fmt.Println(newNote)
			rootNote = newNote
			i++
		}
	} else {
		scale := scaleToNotes(rootNote, notes, *method)
		fmt.Println(scale)
	}
}

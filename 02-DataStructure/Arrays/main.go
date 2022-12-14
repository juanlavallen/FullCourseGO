package main

import "fmt"

func main() {
	var notes [5]int // Default [0 0 0 0 0]
	fmt.Println(notes)

	notes[0] = 10
	notes[1] = 20
	notes[2] = 30
	notes[3] = 40
	notes[4] = 50

	fmt.Println(notes)
	fmt.Println(notes[3])

	characters := [3]string{"Homero", "Marge", "Bart"}
	fmt.Println(characters)

	colors := [...]string{"Red", "Black", "Green", "Orange"}
	fmt.Println(colors, len(colors))

	languages := [...]string{0: "JS", 1: "GO", 2: "Python", 3: "Rust"}
	fmt.Println(languages)

	start := languages[0:2] // => [JS GO] || [:2] = [0:2]
	fmt.Println(start, len(start))

	end := languages[3:] // => [Rust]
	fmt.Println(end, len(end))
}

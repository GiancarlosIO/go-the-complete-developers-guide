package main

import "fmt"

type bot interface {
	getGretting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGretting() string {
	return "Hi there!"
}

func (sb spanishBot) getGretting() string {
	return "Hola!"
}

func printGretting(b bot) {
	fmt.Println(b.getGretting())
}

// this functions has the same logic
// func printGretting(eb englishBot) {
// 	fmt.Println(eb.getGretting())
// }

// func printGretting(sb spanishBot) {
// 	fmt.Println(sb.getGretting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGretting(eb)
	printGretting(sb)
}

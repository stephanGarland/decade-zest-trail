package main

import "fmt"

func main() {
	merchData := merchantDeck()
	scoreData := scoreDeck()
	fmt.Printf("%+v\n", merchData)
	fmt.Printf("%+v\n", scoreData)
}

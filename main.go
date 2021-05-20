package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func shuffle(deck []interface{}) []interface{} {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

func createDecks() ([]byte, []byte) {
	merchData := merchantDeckCreator()
	pointData := PointDeckCreator()
	shuffledMerch := shuffle(merchData)
	shuffledPoint := shuffle(pointData)
	merchJson, _ := json.MarshalIndent(shuffledMerch, "", "    ")
	pointJson, _ := json.MarshalIndent(shuffledPoint, "", "    ")

	return merchJson, pointJson
}

func main() {
	var merchJson, pointJson = createDecks()
	fmt.Print(string(merchJson))
	fmt.Print(string(pointJson))

}

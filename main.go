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

func createDecks() ([]interface{}, []interface{}) {
	merchData := DeckCreator("merchant.json")
	pointData := DeckCreator("point.json")
	shuffledMerch := shuffle(merchData)
	shuffledPoint := shuffle(pointData)

	return shuffledMerch, shuffledPoint
}

func jsonIfyDeck(deck []interface{}) []byte {
	deckJson, _ := json.MarshalIndent(deck, "", "    ")

	return deckJson
}

func main() {
	var merchJson, pointJson = createDecks()
	firstFivePoint := pointJson[:5]
	firstSixMerch := merchJson[:6]
	firstFivePointJson := jsonIfyDeck(firstFivePoint)
	firstSixMerchJson := jsonIfyDeck(firstSixMerch)
	fmt.Print(string(firstFivePointJson))
	fmt.Print(string(firstSixMerchJson))

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func DeckCreator(file string) []interface{} {
	var deck []interface{}
	deckJson, _ := ioutil.ReadFile(file)
	err := json.Unmarshal(deckJson, &deck)
	if err != nil {
		fmt.Println(err)
	}
	return deck
}

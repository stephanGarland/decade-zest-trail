package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func merchantDeckCreator() []interface{} {
	var merchantDeck []interface{}
	merchantJson, _ := ioutil.ReadFile("merchant.json")
	err := json.Unmarshal(merchantJson, &merchantDeck)
	if err != nil {
		fmt.Println(err)
	}
	return merchantDeck
}

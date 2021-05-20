package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func PointDeckCreator() []interface{} {
	var pointDeck []interface{}
	pointJson, _ := ioutil.ReadFile("point.json")
	err := json.Unmarshal(pointJson, &pointDeck)
	if err != nil {
		fmt.Println(err)
	}
	return pointDeck
}

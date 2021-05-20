package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	merchData := merchantDeckCreator()
	pointData := PointDeckCreator()
	merchJson, _ := json.MarshalIndent(merchData, "", "    ")
	pointJson, _ := json.MarshalIndent(pointData, "", "    ")
	fmt.Print(string(merchJson))
	fmt.Print(string(pointJson))
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Table struct {
	FIELD1 string
	FIELD2 string
	FIELD3 string
}

var njia Table

func loadConfig() {
	//takes in the main data file 
	file, err := os.Open("babau.json")
	if err != nil {
		fmt.Println("Cannot open dafile", err)
	}
	decoder := json.NewDecoder(file)
	njia = Table{}
	err = decoder.Decode(&njia)
	if err != nil {
		fmt.Println("Cannot get datfile", err)
	}
}

func main() {
		fmt.Println("Just the start")
}

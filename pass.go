package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type MyJSon []struct {
	Field1	string	`json:"FIELD1"`
	Field2	string	`json:"FIELD2"`
	Field3	string	`json:"FIELD3"`

}

func loadConfig() {
	file, err := os.Open("babau.json")
	if err != nil {
		fmt.Println("Cannot open babau file", err)
	}
	decoder := json.NewDecoder(file)
	myJson := MyJSon{}
	err = decoder.Decode(&myJson)
	if err != nil {
		fmt.Println("Cannot get configuration from file", err)
	} else {
		fmt.Println(myJson)
	}
}

	
func main(){
	loadConfig()
}

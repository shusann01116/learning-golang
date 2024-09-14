package main

import (
	"encoding/json"
	"log"
)

type ID string

type MyJSONType struct {
	ID   ID     `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	serialize()
}

func serialize() {
	v := MyJSONType{
		ID:   "my-id1",
		Name: "John Titor",
	}

	rawstr, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(rawstr))
}

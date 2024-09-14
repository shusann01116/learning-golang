package main

import (
	"encoding/json"
	"log"
	"log/slog"
)

type ID string

type MyJSONType struct {
	ID   ID     `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	serialize()
	deserialize()
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

func deserialize() {
	rawStr := `{
	"id": "oooo",
	"name": "John Doh"
}`
	var result MyJSONType
	if err := json.Unmarshal([]byte(rawStr), &result); err != nil {
		panic(err)
	}

	slog.Info(rawStr)
}

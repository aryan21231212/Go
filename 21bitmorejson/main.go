package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("hello")
	encodeJson()
}

func encodeJson() {
	cources := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "learncodeonline.in", "bot123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "learncodeonline.in", "cba123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(cources, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

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
	// encodeJson()
	decodeJson()
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

func decodeJson() {

	jsonDataFromWeb := []byte(
		`
		 {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "learncodeonline.in",
                "tags": ["web-dev","js"]
        }
		`)

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		var coursera course
		err := json.Unmarshal(jsonDataFromWeb, &coursera)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", coursera)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for k, v := range myonlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}

}

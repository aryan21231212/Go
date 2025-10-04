package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

//middleware - fileter, auth

func (c *Course) isEmpty() bool {
	return c.CourseName == "" && c.CourseId == ""
}

func main() {
	fmt.Println("hello")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API</h1>"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

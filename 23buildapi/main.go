package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	// return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Jane Doe", Website: "janedoe.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Angular", CoursePrice: 299, Author: &Author{Fullname: "Smit Doe", Website: "smitdoe.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Django", CoursePrice: 299, Author: &Author{Fullname: "Kritik Doe", Website: "kritikdoe.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCources).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCource).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCource).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteoneCource).Methods("DELETE")

	//listen to a port
	fmt.Println("Starting server at port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API</h1>"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting one cources")
	w.Header().Set("Content-Type", "application/json")

	// grabing the param
	parama := mux.Vars(r)

	for _, cource := range courses {
		if cource.CourseId == parama["id"] {
			json.NewEncoder(w).Encode(cource)
			return
		}
		json.NewEncoder(w).Encode("no cource found with given id")
	}
}

func createOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one Cource")
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var cource Course
	_ = json.NewDecoder(r.Body).Decode(&cource)
	if cource.isEmpty() {
		json.NewEncoder(w).Encode("no data inside json")
		return
	}

	// generate unique id,string
	rand.Seed(time.Now().UnixNano())
	cource.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, cource)

	json.NewEncoder(w).Encode(cource)
}

func updateOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one Cource")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, cource := range courses {
		if cource.CourseId == params["id"] {
			courses = append(courses[0:index], courses[index+1:]...)
			var cource Course
			json.NewDecoder(r.Body).Decode(&cource)
			cource.CourseId = params["id"]
			courses = append(courses, cource)
			json.NewEncoder(w).Encode(cource)
			return
		} else {
			json.NewEncoder(w).Encode("No cource found with given id")
		}
	}
}

func deleteoneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one cource")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, cource := range courses {
		if cource.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1])
			json.NewEncoder(w).Encode("cource deleted")
			break
		} else {
			json.NewEncoder(w).Encode("No cource found with given id")
		}
	}
}

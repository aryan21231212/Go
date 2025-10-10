package main

import (
	"fmt"
	"github/aryan/mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()

	http.ListenAndServe(":3000", r)
}

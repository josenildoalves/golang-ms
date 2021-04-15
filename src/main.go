package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var counter int

func hello(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	counter++
	fmt.Fprintf(w, "Hello Golang: "+name+" -> "+fmt.Sprint(counter))
	fmt.Println("Hello Golang: " + name + " -> " + fmt.Sprint(counter))
}

func handleRequests() {
	http.HandleFunc("/hello-golang", hello)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	handleRequests()
}

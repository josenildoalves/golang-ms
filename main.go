package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var counter int

func hello(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	counter++
	fmt.Fprintf(w, "Hello Golang: "+name+" -> "+fmt.Sprint(counter))
	fmt.Println(currentTime.Format("01-02-2006 15:04:05") + " - Hello Golang: " + name + " -> " + fmt.Sprint(counter))
}

func handleRequests() {
	http.HandleFunc("/hello-golang", hello)
	fmt.Println("Golang Microservice started at port 9001")
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func main() {
	handleRequests()
}

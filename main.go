package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var counter int

func helloWorld(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	hostName, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	counter++
	fmt.Fprintf(w, "Hello World Golang: "+hostName+" -> "+fmt.Sprint(counter))
	fmt.Println(currentTime.Format("01-02-2006 15:04:05") + " - Hello Golang: " + hostName + " -> " + fmt.Sprint(counter))
}

func handleRequests() {
	http.HandleFunc("/hello-golang", helloWorld)
	fmt.Println("Golang Microservice started at port 9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	handleRequests()
}

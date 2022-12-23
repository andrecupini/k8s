package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	// http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I'm %s. I'm %s", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s ", err)
	}

	fmt.Fprintf(w, "My ffamily: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {

	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User: %s Password: %s", user, pass)
}

// func Healthz(w http.ResponseWriter, r *http.Request) {

// 	duration := time.Since(startedAt)

// 	if duration.Seconds() < 15 || duration.Seconds() > 40 {
// 		w.WriteHeader(500)
// 		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
// 	} else {
// 		w.WriteHeader(200)
// 		w.Write([]byte(fmt.Sprintf("ok")))
// 	}
// }

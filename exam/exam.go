package main

import (
	"fmt"
	"net/http"
)

type Task struct {
	ID int
	Description int
	Done bool
}

var task []Task

func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8080", nil)
}

func list(rw http.ResponseWriter, _ *http.Request) {
	for i, v := range task {
		fmt.Println(i, v)
	}
	rw.WriteHeader(http.StatusOK)
}

func done(d http.ResponseWriter, _ *http.Request) {
}
func add(a http.ResponseWriter, _ *http.Request) {
}

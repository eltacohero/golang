package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
//define Task type structure
type Task struct {
	ID int
	Description int
	Done bool
}
//define task var as a slice of Task
var task []Task
//main function
func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8080", nil)
}
//list function
func list(rw http.ResponseWriter, _ *http.Request) {
	for i, v := range task {
		if Done == false {
			fmt.Println(i, v)
		}
	}
	rw.WriteHeader(http.StatusOK)
}
//done function
func done(rw http.ResponseWriter, _ *http.Request) {
}
//add function
func add(rw http.ResponseWriter, _ *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body:%v", err)
		http.Error(
		rw,
		"can't read body", http.StatusBadRequest,
		)
		return
	}
}

package main

import (
	"fmt"
	"encoding/json"
)

type Users struct {
	ID int
	Login string `json:"Username"`
	Password string
} 

func main() {
	group := Users{
		ID: 1,
		Login: "Paul",
		Password: "pass123",
	}
	b, err := json.Marshal(group)
	if err != nil{
		fmt.Println("error", err)
	}
	fmt.Println(string(b))
}

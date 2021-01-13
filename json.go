package main

import (
	"fmt"
	"encoding/json"
)

type Users struct {
	Login string
	Password string
} 

func main() {
	group := Users{
		Login: "Paul",
		Password: "pass123",
	}
	b, err := json.Marshal(group)
	if err != nil{
		fmt.Println("error", err)
	}
	fmt.Println(string(b))
}

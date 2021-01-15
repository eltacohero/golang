package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	ID int
	Login string `json:"Username"`
	Password string
}

type Users []User

func main() {
	users := Users{
	User{
		ID: 1,
		Login: "Paul",
		Password: "pass123",
	},
	User{
		ID: 2,
		Login: "matm",
		Password: "123456",
	},
	User{
		ID: 3,
		Login: "fake44",
		Password: "azerty",
	},
}
	b, err := json.Marshal(users)
	if err != nil{
		fmt.Println("error", err)
	}
	fmt.Println(string(b))
}

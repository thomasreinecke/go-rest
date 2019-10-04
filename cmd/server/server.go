package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thomasreinecke/go-rest/internal/user"
)

func main() {
	fmt.Println("Hello, world.")
	//fmt.Println(stringutil.Reverse("!oG ,olleH"))

	nu := user.NewUser{
		Name:     "Thomas Reinecke",
		Email:    "reinecke.fox@googlemail.com",
		Password: "password",
	}
	fmt.Println("standard serialized representation: ", nu)

	nuJSON, _ := json.Marshal(nu)
	fmt.Println("JSON serialized representation: ", string(nuJSON))

	u, err := user.Create(nu, time.Now())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
}

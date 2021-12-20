package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	user := User{
		Email:    "feamo@ya.ru",
		Password: "SSSSS",
	}

	userBytes, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/user", "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
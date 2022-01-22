package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//func userGet() {
//	resp, err := http.Get("http://localhost:8080/user/1")
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(string(body))
//}

//func userGetList() {
//	resp, err := http.Get("http://localhost:8080/user")
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
//	fmt.Println(string(body))
//
//	var result []User
//	err = json.Unmarshal(body, &result)
//	if err != nil { // Parse []byte to the go struct pointer
//		fmt.Println("Can not unmarshal JSON")
//	}
//	fmt.Println(PrettyPrint(result))
//}
//
//// PrettyPrint to print struct in a readable way
//func PrettyPrint(i interface{}) string {
//	s, _ := json.MarshalIndent(i, "", "\t")
//	return string(s)
//}

//func userCreate() {
//	user := User{
//		Email:    "11dcgh@ya.ru",
//		Password: "SDdgd",
//	}
//
//	userBytes, err := json.Marshal(&user)
//	if err != nil {
//		panic(err)
//	}
//
//	resp, err := http.Post("http://localhost:8080/user", "application/json", bytes.NewBuffer(userBytes))
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(string(body))
//}

//func UserUpdate() {
//	user := User{
//		ID:       3,
//		Email:    "feamo@ya.ru",
//		Password: "s888",
//	}
//
//	// initialize http client
//	client := &http.Client{}
//
//	// marshal User to json
//	userBytes, err := json.Marshal(user)
//	if err != nil {
//		panic(err)
//	}
//
//	// set the HTTP method, url, and request body
//	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/user", bytes.NewBuffer(userBytes))
//	if err != nil {
//		panic(err)
//	}
//
//	// set the request header Content-Type for json
//	resp, err := client.Do(req)
//	if err != nil {
//		panic(err)
//	}
//
//defer resp.Body.Close()
//
//	fmt.Println(resp.StatusCode)
//}

func main() {
	user := User{
		ID: 1,
	}

	// initialize http client
	client := &http.Client{}

	// marshal User to json
	userBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/user/1", bytes.NewBuffer(userBytes))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}

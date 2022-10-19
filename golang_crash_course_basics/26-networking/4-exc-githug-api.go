package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	NumRepos int    `json:"public_repos"`
}

func getUserInfo(login string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(res.Status)
	}

	user := User{Login: login}
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil

}

func main() {
	res, err := getUserInfo("pkro")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", *res)
}

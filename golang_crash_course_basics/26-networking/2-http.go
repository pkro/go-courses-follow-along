package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call url")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("------------------------------------")

	// POST request
	type Job struct {
		User   string `json:"user"`
		Action string `json:"action"`
		Count  int    `json:"count"`
	}

	job := &Job{
		User:   "pkro",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: can't call url - %v", err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't post to url")
	}
	io.Copy(os.Stdout, resp.Body)
}

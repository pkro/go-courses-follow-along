package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

var data = `
{
	"user": "pkro",
	"type": "deposit",
	"amount": 123.4
}
`

func main() {
	rdr := strings.NewReader(data) // simulate file / socket

	// decode request
	dec := json.NewDecoder(rdr)

	var req Request

	// param must be pointer
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("error: can't decode %s", err)
	}

	fmt.Printf("got %v\n", req)

	// create rerespsponse
	prevBalance := 1_000_000.0 // simulate loaded from db
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// Encode response

	enc := json.NewEncoder(os.Stdout) // passes to stdout, so no println necessary

	if err := enc.Encode(resp); err != nil {
		log.Fatalf("couldn't decode, %s", err)
	}
}

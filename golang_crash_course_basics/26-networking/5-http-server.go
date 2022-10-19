package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	// step 1: decode and validate
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}

	if err := dec.Decode(req); err != nil {
		log.Printf("error: bad JSON: %s", err)
		http.Error(w, "bad JSON", http.StatusBadRequest)
		return
	}

	if !strings.Contains("+-*/", req.Op) {
		log.Printf("error: bad operator: %q", req.Op)
		http.Error(w, "unknown operator", http.StatusBadRequest)
	}

	// step 2: work
	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		resp.Result = req.Left / req.Right
	default:
		resp.Error = fmt.Sprintf("unknon operation: %s", req.Op)
	}

	// step 3: Encode result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("can't encode %v - %s", resp, err)
	}
}

func main() {

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/math", mathHandler)

	addr := ":8030"
	log.Printf("server running on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

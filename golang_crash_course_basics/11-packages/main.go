package main

import (
	"example.com/packages/strutil"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println(strutil.Reverse("pkro"))
	fmt.Println(strutil.ReverseAgain("pkro"))

	r := mux.NewRouter()
	http.ListenAndServe(":9000", r)
}

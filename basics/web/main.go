package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just hit the slash root route")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3001", nil)
}

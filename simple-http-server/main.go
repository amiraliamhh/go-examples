package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Homepage!")
}

func main() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}

	// run `curl http://127.0.0.1:8000/` to test
}

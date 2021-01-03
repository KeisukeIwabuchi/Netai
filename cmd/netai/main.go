package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("寝たい")

	fs := http.FileServer(http.Dir("web"))

	http.Handle("/", fs)

	http.ListenAndServe(":8888", nil)
}

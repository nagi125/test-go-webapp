package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
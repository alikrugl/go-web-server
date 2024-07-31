package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8080", mux)
    if err != nil {
        panic(err)
    }
}

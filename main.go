package main

import "net/http"

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}

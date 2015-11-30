package main

import (
	"net/http"

	"padpad/server"
)

// This example demonstrates a trivial echo server.
func main() {
	server := server.NewServer("/handler")
	go server.Listen()

	http.Handle("/", http.FileServer(http.Dir("pub")))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

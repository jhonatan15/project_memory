package main

import (
	"net/http"
)

func main() {
	// Route to serve static files of the frontend
	fs := http.FileServer(http.Dir("frontend/build"))
	http.Handle("/", fs)

	// Iniciar el servidor
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

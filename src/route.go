package main

import (
	"fmt"
	"log"
	"net/http"
)

// Permet de faire les routes vers les differentes pages.
func route(h1, h2, h3 http.HandlerFunc) {
	http.HandleFunc("/", h1)
	http.HandleFunc("/about", h2)
	http.HandleFunc("/concerts", h3)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	ClearConsole()
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

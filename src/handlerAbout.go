package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//Permet d'afficher la page about
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	page := "about"
	data := struct {
		Page string
	}{
		Page: page,
	}
	filePath := "./static/html/about.html"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

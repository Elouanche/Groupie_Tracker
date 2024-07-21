package main

import (
	"net/http"
)

// Permet de charger la page html que l'on souhaite
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	err := tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	} 
}

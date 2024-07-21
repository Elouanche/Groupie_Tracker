package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var (
	tmpl *template.Template
	page = " "
)

// fonction principal du code
func main() {
	help()
	route(
		handlerArtistes,
		handlerAbout,
		handlerRelation,
	)
}

// Permets d'initialiser toutes les pages html
func init() {
	var err error
	tmpl, err = template.ParseGlob("./static/html/*.html")
	if err != nil {
		panic(err)
	}
}

func handleFormData(r *http.Request) {
	text := strings.ToLower(r.FormValue("String"))
	fmt.Println(text)
}

func handleHTTPError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

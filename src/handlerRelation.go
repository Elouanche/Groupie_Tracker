package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

// permets d'envoyer les données de l'api
func handlerRelation(w http.ResponseWriter, r *http.Request) {
	artistID := r.FormValue("artistID")
	artistName := r.FormValue("artistName")
	artistImage := r.FormValue("artistImage")
	artistYearStarted := r.FormValue("artistYearStarted")
	artistFirstAlbum := r.FormValue("artistFirstAlbum")
	artistMembers := strings.Split(r.FormValue("artistMembers"), ",")

	apiURLs := map[string]string{
		"api": "https://groupietrackers.herokuapp.com/api/relation",
	}

	apiData, err := RelationResponse(apiURLs, artistID)
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	data := struct {
		API               map[string][]relation
		ArtistName        string
		ArtistImage       string
		ArtistYearStarted string
		ArtistFirstAlbum  string
		ArtistMembers     []string
	}{
		API:               apiData,
		ArtistName:        artistName,
		ArtistImage:       artistImage,
		ArtistYearStarted: artistYearStarted,
		ArtistFirstAlbum:  artistFirstAlbum,
		ArtistMembers:     artistMembers,
	}

	tmpl, err := template.ParseFiles("./static/html/Location.html")
	if err != nil {
		handleHTTPError(w, fmt.Errorf("Erreur lors de l'analyse du template HTML: %v", err))
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		handleHTTPError(w, fmt.Errorf("Erreur lors de l'exécution du template HTML: %v", err))
		return
	}
}


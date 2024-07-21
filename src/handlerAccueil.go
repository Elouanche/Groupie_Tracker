package main

import (
	"net/http"
)

// Permets de charger et d'envoyer les données de l'api
func handlerArtistes(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	data, err := fetchDataFromAPI(apiURL)
	if err != nil {
		http.Error(w, "Fail dans le chargement des données :", http.StatusInternalServerError)
		return
	}

	apiData := map[string]interface{}{
		"artists": data,
	}

	err = tmpl.ExecuteTemplate(w, "index.html", apiData)
	if err != nil {
		http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		return
	}
}

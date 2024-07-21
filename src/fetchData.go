package main

import (
	"encoding/json"
	"net/http"
)

// Permets de charger les données de l'api que l'on souhaite charger
func fetchDataFromAPI(apiURL string) (interface{}, error) {
	// Envoi d'une requête HTTP GET pour récupérer les données de l'API
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data interface{}
	// Décode les données JSON de la réponse HTTP dans la variable data
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

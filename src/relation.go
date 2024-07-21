package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RelationResponse(apiURLs map[string]string, artistID string) (map[string][]relation, error) {
	// Création d'une structure pour stocker les données de chaque API
	apiData := make(map[string][]relation)

	// Parcourir les URLs des APIs
	for key, url := range apiURLs {
		// Effectuer une requête HTTP GET vers l'API
		resp, err := http.Get(url)
		if err != nil {
			// Si une erreur se produit, renvoyer une réponse d'erreur HTTP 500
			errorMsg := fmt.Sprintf("Erreur lors de la requête GET pour %s: %v", key, err)
			fmt.Println(errorMsg)
			return nil, fmt.Errorf(errorMsg)
		}
		defer resp.Body.Close() // Fermer le corps de la réponse après la fin de la fonction

		// Décodez la réponse JSON de l'API
		var responseData APIRelation
		err = json.NewDecoder(resp.Body).Decode(&responseData)
		if err != nil {
			// Si une erreur se produit lors du décodage JSON, renvoyer une réponse d'erreur HTTP 500
			errorMsg := fmt.Sprintf("Erreur lors du décodage de la réponse JSON de %s: %v", key, err)
			fmt.Println(errorMsg)
			return nil, fmt.Errorf(errorMsg)
		}

		// Parcourir les données de l'API et les organiser correctement
		var relations []relation
		for _, item := range responseData.Index {
			// Vérifier si l'ID de l'artiste correspond à celui fourni
			if strconv.Itoa(item.ID) != artistID {
				continue
			}

			// Créer une nouvelle relation
			rel := relation{
				ID: item.ID,
			}
			// Parcourir les dates et lieux de l'événement
			for date, locations := range item.DatesLocations {
				for _, loc := range locations {
					// Ajouter l'événement à la relation
					rel.Events = append(rel.Events, Event{
						Date:     date,
						Location: loc,
					})
				}
			}
			// Ajouter la relation à la liste des relations
			relations = append(relations, rel)
		}

		// Ajouter les relations à la map des données de l'API
		apiData[key] = relations
	}

	return apiData, nil
}

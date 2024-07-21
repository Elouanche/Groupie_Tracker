package main

// Structure pour les informations sur un artiste
type Artist struct {
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	YearStarted int      `json:"year_started"`
	FirstAlbum  string   `json:"first_album"`
	Members     []string `json:"members"`
}

//structure pour l'api relation

type location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type APILocation struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Event struct {
	Artist   string `json:"name"`
	Date     string `json:"date"`
	Location string `json:"location"`
}

type relation struct {
	Name   string  `json:"name"`
	ID     int     `json:"id"`
	Events []Event `json:"events"`
}

type APIRelation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}


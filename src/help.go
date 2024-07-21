package main

import (
	"flag"
	"fmt"
	"os"
)

// cette fonction permets d'afficher les aides essentiels pour le jeu
func help() {
	helpFlag := flag.Bool("h", false, "Affiche les explications")

	flag.Parse()

	// Vérification de l'option -h
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}
}

// Fonction qui permets de clear le cmd

// fonction qui permets d'afficher les aides
func printHelp() {
	ClearConsole()
	fmt.Println("Bienvenue sur le groupie tracker. ")
	fmt.Println("\nAfin de lancer le seveur vous devez dans un premier temps écrire dans votre cmd :\n make build puis make run ou .\\groupie-tracker.exe ")
	fmt.Println("\nEt dans un deuxieme allé sur la bare de recherche de votre navigateur et taper 'localhost:8080' OU cliquer sur le lien qui s'affiche dans le cmd ")
	fmt.Println("\nConditions d'utilisations :")
	fmt.Println("Après avoir fait make build il ne faut pas modifier les fichiers, cela peut poser problème.")
	fmt.Println("N'ajouter rien afin de ne pas casser le fonctionnement du site")
	fmt.Println("Avant de modifier le programme, faites un 'make clean' ou  `make restart` afin de clean, build et run après avoir fais des modifications.")
	fmt.Println("Amusez-vous bien et bonne chance avec toutes ces musiques à porté de main ! 🎉✨")
}

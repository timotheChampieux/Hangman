package recupmot

import (
	"bufio"
	"fmt"
	"os"
)

func recup() []string {
	file, err := os.Open("mot.txt") // on ouvre le fichier mot.go
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // lecture du fichier ligne par ligne
	var mots []string
	var i int

	for scanner.Scan() {
		ligne := scanner.Text() // Récupérer la ligne actuelle
		mots[i] = ligne         // Ajouter les mots à la liste
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}

	return mots // retourne la liste des mots récupérés
}

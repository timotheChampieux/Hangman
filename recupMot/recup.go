package recupmot

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func Recup() string {

	file, _ := os.Open(".\\recupMot\\mot.txt") // on ouvre le fichier mot.go
	defer file.Close()                         // garantie que e fichier sera fermé une fois que la fonction aura terminé son exécution
	scanner := bufio.NewScanner(file)          // lecture du fichier ligne par ligne

	var mots []string
	var i int

	for scanner.Scan() {
		ligne := scanner.Text()    // Récupérer la ligne actuelle
		mots = append(mots, ligne) // Ajouter les mots à la liste
		i++
	}

	rand.Seed(time.Now().UnixNano())
	max := len(mots)
	indexMotAleatoire := rand.Intn(max)
	return mots[indexMotAleatoire] // retourne le mot aleatoire a trouver
}

package main

import (
	"fmt"
	"hangman/affichage"
	recupmot "hangman/recupMot"
)

func main() {
	motAleatoire := recupmot.Recup()
	fmt.Println(motAleatoire)
	affichage.Debut(motAleatoire)
}

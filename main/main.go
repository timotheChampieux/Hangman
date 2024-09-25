package main

import (
	"fmt"
	"hangman/affichage"
	"hangman/jeu"
	recupmot "hangman/recupMot"
)

func main() {
	motAleatoire := recupmot.Recup()
	fmt.Println(motAleatoire)
	motMasqer := affichage.Debut(motAleatoire)
	jeu.Jeu(motAleatoire, motMasqer)
}

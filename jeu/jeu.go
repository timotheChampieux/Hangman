package jeu

import (
	"fmt"
	"hangman/affichage"
	"strings"
)

func demanderElement(slice *[]string) string {
	for {
		fmt.Print("Entrez une lettre ou un mot: ")
		var lettre string
		fmt.Scanln(&lettre)

		if !elementDansSlice(lettre, *slice) {
			*slice = append(*slice, lettre)
			return lettre
		} else {
			fmt.Println("Vous avez déja choisi cette lettre ou ce mot veuillez ressayer.")
		}
	}
}

func elementDansSlice(element string, slice []string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func Jeu(mot string, motMasque []string) {
	essaie := 8
	lettreDejaPropose := []string{}
	var reussitte int
	if len(motMasque) <= 5 {
		reussitte = len(motMasque) - 1
	} else {
		reussitte = len(motMasque) - 2
	}
	for essaie > 1 && reussitte > 0 {

		lettre := demanderElement(&lettreDejaPropose)

		fmt.Println(lettreDejaPropose)
		indexDebut := strings.Index(mot, lettre)
		if indexDebut == -1 {
			if len(lettre) < 2 {
				essaie--
				affichage.AfficherPendu(essaie)
			} else {
				essaie -= 1
				affichage.AfficherPendu(essaie)
				essaie -= 1
				affichage.AfficherPendu(essaie)
			}

			fmt.Printf("La lettre que vous avez entré n'est pas contenue dans le mot, il vous reste %v vies \n", essaie)
		} else {
			fmt.Printf("Vous avez trouvé, il vous reste %v vies \n", essaie-1)
			for index, _ := range lettre {
				fmt.Print("\n", reussitte)
				if motMasque[indexDebut] != string(lettre[index]) {
					motMasque[indexDebut] = string(lettre[index])
					reussitte--
				}

				indexDebut++

			}
		}

		for i := 0; i < len(motMasque); i++ {
			fmt.Printf("%v ", motMasque[i])
		}
	}
	if reussitte <= 0 {
		fmt.Println("\nVous avez gagné !!")
	} else {
		fmt.Println("\nVous avez perdu...")

	}
}

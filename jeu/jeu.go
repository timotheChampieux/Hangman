package jeu

import (
	"fmt"
	"hangman/affichage"
)

func demanderElement(slice *[]string) string {
	for {
		fmt.Print("\nElement déja choisi : ")
		for _, char := range *slice {
			fmt.Print(char, "  ")
		}
		fmt.Print("\n\033[32mEntrez une lettre ou un mot:\033[0m \n")
		var lettre string
		fmt.Scanln(&lettre)

		if !elementDansSlice(lettre, *slice) {
			*slice = append(*slice, lettre)
			return lettre
		} else {
			fmt.Println("\n\033[31mVous avez déja choisi cette lettre ou ce mot veuillez ressayer.\033[0m\n")
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
		fmt.Println(reussitte)

		lettre := demanderElement(&lettreDejaPropose)

		if len(lettre) == 1 {
			var count int
			for i := 0; i < len(mot); i++ {
				if lettre == string(mot[i]) && lettre != motMasque[i] {
					motMasque[i] = lettre
					reussitte--
					count++
				}
			}
			if count == 0 {
				essaie--
				fmt.Printf("\n\033[31mLa lettre que vous avez entré n'est pas contenue dans le mot, il vous reste %v vies\033[0m \n", essaie)
				affichage.AfficherPendu(essaie)
			} else {
				for i := 0; i < len(motMasque); i++ {
					fmt.Printf("%v ", motMasque[i])
				}
			}
		} else {
			if lettre == mot {
				for i := 0; i < len(mot); i++ {
					motMasque[i] = string(lettre[i])
				}
				for i := 0; i < len(motMasque); i++ {
					fmt.Printf("%v ", motMasque[i])
				}
				reussitte = 0
			} else {
				essaie -= 2
				affichage.AfficherPendu(essaie)
			}

		}
	}
	if reussitte <= 0 {
		fmt.Println("\n\033[32mVous avez gagné !!\033[0m")
	} else {
		fmt.Println("\n\033[31mVous avez perdu...\033[0m")

	}
}

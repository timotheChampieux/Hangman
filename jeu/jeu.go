package jeu

import (
	"fmt"
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
	essaie := 7
	lettreDejaPropose := []string{}
	for essaie > 0 {
		lettre := demanderElement(&lettreDejaPropose)
		c := false
		fmt.Println(lettreDejaPropose)
		for i := 0; i < len(mot); i++ {
			if string(mot[i]) == lettre {
				motMasque[i] = lettre
				c = true
			}
		}
		for i := 0; i < len(motMasque); i++ {
			fmt.Printf("%v ", motMasque[i])
		}
		if !c {
			essaie--
		}
	}
}

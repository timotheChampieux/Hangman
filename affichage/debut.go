package affichage

import (
	"fmt"
	"math/rand"
	"time"
)

func Debut(mot string) []string {
	motMasque := []string{}
	for i := 0; i < len(mot); i++ {
		motMasque = append(motMasque, "_")
	}
	if len(mot) < 5 {
		rand.Seed(time.Now().UnixNano())
		max := len(motMasque)
		indexLettreAletoire := rand.Intn(max)
		motMasque[indexLettreAletoire] = string(mot[indexLettreAletoire])
	}
	if len(mot) > 4 {
		max := len(motMasque)
		indexLettreAletoire := rand.Intn(max)
		indexLettreAletoire1 := rand.Intn(max)
		motMasque[indexLettreAletoire] = string(mot[indexLettreAletoire])
		motMasque[indexLettreAletoire1] = string(mot[indexLettreAletoire1])
	}
	for i := 0; i < len(motMasque); i++ {
		fmt.Printf("%v ", motMasque[i])
	}
	return motMasque
}

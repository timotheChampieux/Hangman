package son

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

// Variables globales
var player *oto.Player
var context *oto.Context
var once sync.Once

// Fonction pour initialiser le contexte oto (appelée une seule fois)
func initOtoContext() error {
	var err error
	once.Do(func() { // Exécuter une seule fois
		// Ajuster les paramètres audio : 44100 Hz, 2 canaux (stéréo), 16 bits
		context, err = oto.NewContext(48000, 2, 2, 11192) // Taille du buffer augmentée
	})
	return err
}

// Fonction pour jouer un fichier audio MP3
func JouerSon(cheminFichier string) error {
	StopSon() // Arrêter tout son en cours de lecture

	// Initialiser le contexte oto (appelé une seule fois)
	err := initOtoContext()
	if err != nil {
		return fmt.Errorf("erreur d'initialisation du contexte oto: %v", err)
	}

	// Ouvre le fichier audio
	f, err := os.Open(cheminFichier)
	if err != nil {
		return fmt.Errorf("erreur d'ouverture du fichier: %v", err)
	}
	defer f.Close()

	// Décoder le fichier MP3
	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		return fmt.Errorf("erreur de décodage MP3: %v", err)
	}

	// Crée le player oto (à partir du contexte initialisé une seule fois)
	player = context.NewPlayer()
	defer player.Close()

	// Lit et envoie les données audio au player
	if _, err := io.Copy(player, decoder); err != nil {
		return fmt.Errorf("erreur de lecture du fichier audio: %v", err)
	}

	return nil
}

// Fonction pour arrêter la lecture du son
func StopSon() {
	if player != nil {
		player.Close() // Arrête et libère le lecteur en cours
		player = nil
	}
}

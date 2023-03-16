package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

// package permettant de lire le dossier words afin de selctionner un mot de et l'introduire a une fonction
var words = make([]string, 0, 200)

func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func PickWord() string {
	//selectionne un mot aleatoirement
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}

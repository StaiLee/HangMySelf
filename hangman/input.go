package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// partie corespondant a la saisie clavier afin de recuperer la lettre a l'aide dun scan du joueur
// notre variable reader va nous premettre de lire la proposition du joueur
var reader = bufio.NewReader(os.Stdin)

// ReadGuess ne prendra pas de parametre mais va seulement nous return une proposition
func ReadGuess() (guess string, err error) {
	//initialisation d'une boucle while tres pratique dans notre situation
	valid := false
	for !valid {
		fmt.Print("What is your letter ? ")
		//le scan va lire jusqua que le joueur utilise la touche entree qui se materialise donc par un /n
		guess, err = reader.ReadString('\n')
		//filtrages des erreurs dans la boucle
		if err != nil {
			return "", err
		}
		//enlevement des espaces afin de recuperer un seul caractere et non pas l'espace qui va avec
		guess = strings.TrimSpace(guess)
		// si le joueur va utiliser plus d'une lettre, un message derreur va apparaitre, la boucle continue et il devra retapper une lettre
		if len(guess) != 1 {
			fmt.Printf("Invalid letter size. letter=%v, len=%v", guess, len(guess))
			continue
		}
		//donc si c'est bon valid egal = true fin de la boucle et on obtiens donc notre return
		valid = true
	}
	return guess, nil
}

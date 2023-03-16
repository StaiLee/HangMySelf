package hangman

import (
	"strings"
)

type Game struct {
	State        string   // etat de le partie
	Letters      []string // Lettre dans le mot a utiliser
	FoundLetters []string // bon choix
	UsedLetters  []string // lettres utilisé
	TurnsLeft    int      // tentatives restantes
}

// new va preparer notre jeu et renvoyer un pointeur de game pour ne pas ompacter directement la struct
func New(turns int, word string) (*Game, error) {
	//preparation des differentes variables que l'on va utiliser plus tard

	//letters va split le "word" dans une chaine de caractere vide, que l'on a mis en majuscule au prealable
	letters := strings.Split(strings.ToUpper(word), "")
	// Found va etre les lettres trouver stocker dans une variable intermediaire il faut forcement qu'il est la meme taille des lettres que l'on a trouver
	found := make([]string, len(letters))
	//les lettres trouver vont etre initialiser par la valeur _ donc ils vont etre cacher a l'ecran
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	// g va etre un pointeur de la struct game donc modifier le "jeu" sans passer directement par game
	g := &Game{
		State:        "",         //chaine vide initialisation par default
		Letters:      letters,    // les lettres dans le mot
		FoundLetters: found,      // les lettres trouvés
		UsedLetters:  []string{}, //un slice vide se slice grandira vue qu'il n'y a pas de proposition par default
		TurnsLeft:    turns,      // les nombres de tours
	}
	return g, nil // simple return de g donc d'un pointeur de game
}

// logique du jeu initialisation du jeu
func (g *Game) MakeAGuess(guess string) {
	//pour etre sur que nos comparaison fonctionne bien, on met toutes les lettres en majuscules
	guess = strings.ToUpper(guess)

	switch g.State {
	case "won", "lost":
		return
	}
	// associe l'etat au state afin qu'a chaque tours de boucle l'etat du jeu se dessine

	//donc ici already guess
	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
		//ici goodguess et on reveal la lettre qui est un acutellement un _
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		//cas de victoire
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		//cas de defaite
		g.State = "badGuess"
		g.LoseTurn(guess)

		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

// fonction qui va comparer les lettres trouver et renvoyer si cest vraie ou faux
func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

// la fonction reveal letter va modifier le _ et afficher la bonne lettre
func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters { //on va parcourir toutes les lettres trouvés
		if l == guess { //et si la lettre trouver et egal a une lettre du mot on redefinis l'index afin d'afficher la lettre
			g.FoundLetters[i] = guess
		}
	}
}

// le pendu commence et se termine lorsque l'on a plus de tour loseturn va decrementer
func (g *Game) LoseTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

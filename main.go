package main

import (
	"fmt"
	"hangmyselfw/hangman"
	"os"
)

// boucle de jeu, mise en place des differentes fonctions crees
func main() {
	// mise en place du dictionnaire
	err := hangman.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load liste: %v\n", err)
		os.Exit(1) //si on a pas pu lire la saisie clavier on va quitter le programme
	}
	//mise en place de la loop
	g, err := hangman.New(10, hangman.PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}
	//initialise une boucle infinis
	// a chaque tour de boucle notre partie va s'actualiser
	hangman.DrawWelcome()
	//guess est vide pour demarrer
	guess := ""
	for {
		hangman.Draw(g, guess)
		//switch marche aussi avec if pour sortir de la boucle si le joueur a perdu ou bien gagner
		switch g.State {
		case "won", "lost":
			os.Exit(0) //donc on sors avec un code erreur 0 donc pas de probleme
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		//on appel la fonction make a guess qui va executer les choix avec en parametre guess qui est pour l'instant une chaine vide
		guess = l
		g.MakeAGuess(guess)
	}
}

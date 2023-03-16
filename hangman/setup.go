package hangman

import (
	"fmt"
	"hangmyselfw/art"
)

// package permettant de "dessiner" l'etat de la partie en effectuant des switch case et d'incorporer l'ascii art au sein du jeu

// fonction de depart
func DrawWelcome() {
	art.AsciiArt()
}

// fonction Draw est le pillier elle va appeller les petites fonctions qui vont dessiner le jeu
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft) // affichage du hangman selon le tours
	drawState(g, guess)    // etat de la partie ou est on dans la partie
	fmt.Println()
}

// ETAT DE LA PARTIE AVEC ASCII ART
func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		`
		art.One()
	case 1:
		draw = `
		`
		art.Two()
	case 2:
		draw = `
		`
		art.Three()
	case 3:
		draw = `
		`
		art.Four()
	case 4:
		draw = `
		`
		art.Five()
	case 5:
		draw = `
		`
		art.Six()
	case 6:
		draw = `
		`
		art.Seven()
	case 7:
		draw = `
		`
		art.Eight()
	case 8:
		draw = `
		`
		art.Nine()
	case 9:
		draw = `
        `
		art.Ten()
	case 10:
		draw = `
        `
	}
	fmt.Println(draw)
}

// fonction qui va boucler sur les lettres trouvers pour nous aider aux niveaux de l'affichage
func drawLetters(g []string) {
	for _, c := range g {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}

// donc drawLetters va boucler sur les lettres trouvers et va donc utiliser found et used pour l'affichage
func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)
	// maitenant on va switch sur g.state qui va nous indiquer si l'etat de la lettre donc utiliser mauvaise....
	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess !")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word ", guess)
	case "lost":
		fmt.Print("You lost ! The word was: ")
		drawLetters(g.Letters)
		art.Zero()
	case "won":
		fmt.Print("you won ! The word was: ")
		drawLetters(g.Letters)
		art.Victory()
	}
	fmt.Println()
}

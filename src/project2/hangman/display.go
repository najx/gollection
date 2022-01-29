package hangman

import (
	"fmt"
)

func DrawWelcome() {
	fmt.Println(`
	▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
	██░██░█░▄▄▀█░▄▄▀█░▄▄▄█░▄▀▄░█░▄▄▀█░▄▄▀██
	██░▄▄░█░▀▀░█░██░█░█▄▀█░█▄█░█░▀▀░█░██░██
	██░██░█▄██▄█▄██▄█▄▄▄▄█▄███▄█▄██▄█▄██▄██
	▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
	fmt.Println()
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		[X] [X] [X] [X] [X] [X] [X] [X] :-(
    `
	case 1:
		draw = `
		[X] [X] [X] [X] [X] [X] [X] [ ] :-/
		`
	case 2:
		draw = `
		[X] [X] [X] [X] [X] [X] [ ] [ ] :-/
		`
	case 3:
		draw = `
		[X] [X] [X] [X] [X] [ ] [ ] [ ] :-/
		`
	case 4:
		draw = `
		[X] [X] [X] [X] [ ] [ ] [ ] [ ] :-/
		`
	case 5:
		draw = `
		[X] [X] [X] [ ] [ ] [ ] [ ] [ ] :-/
		`
	case 6:
		draw = `
		[X] [X] [ ] [ ] [ ] [ ] [ ] [ ] :-/
		`
	case 7:
		draw = `
		[X] [ ] [ ] [ ] [ ] [ ] [ ] [ ] :-/
    `
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawLetters(g []string) {
	for _, c := range g {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!")
	case "alreadyGuessed":
		fmt.Println("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Println("Bad guess, '%s' is not in the word\n", guess)
	case "lost":
		fmt.Print("You lost :(! The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON! The word was: ")
		drawLetters(g.Letters)
	}
	fmt.Println()
}

# Blackjack

This program is a simple implementation of the card game Blackjack in Go. It allows a player to play against a dealer and determine the winner based on the score of the hands.
Running the program

To run the program, make sure you have Go installed and configured on your machine. Then, navigate to the directory containing the program and run the following command:

```go
go run main.go
```

## Playing the game

When the program starts, the player and dealer will be dealt two cards each. The player's hand and score will be displayed, and only the first card of the dealer's hand will be shown.

The player can then choose to "hit" and be dealt another card, or "stand" to keep their current hand. The player can continue to hit until they have a score of 21 or more, at which point the game will proceed to the dealer's turn.

If the player's score exceeds 21, they will lose the game. If the player chooses to stand, the dealer's full hand and score will be revealed and the dealer will continue to hit until they have a score of 17 or more.

If the dealer's score exceeds 21, the player wins. Otherwise, the winner will be determined by comparing the player's and dealer's scores. If the scores are equal, the game will end in a tie.

    Note

    . The game is implemented for single deck.
    . The max score for the game is 21.
    . When the deck ends, it will be re-initialized and shuffled again.

<p align="center"><img src="https://user-images.githubusercontent.com/60783263/214822283-8a90907c-d83e-45f7-b525-ae20cf62c495.png" /></p>

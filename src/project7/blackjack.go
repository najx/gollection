package main

import (
  "fmt"
  "math/rand"
  "time"
)

const (
  numDecks = 1
  maxScore = 21
)

type card struct {
  suit string
  rank string
}

var deck []card

func initDeck() {
  suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
  ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
  for i := 0; i < numDecks; i++ {
    for _, suit := range suits {
      for _, rank := range ranks {
        deck = append(deck, card{suit, rank})
      }
    }
  }
}

func shuffleDeck() {
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}

func drawCard() card {
  if len(deck) == 0 {
    initDeck()
    shuffleDeck()
  }
  c := deck[0]
  deck = deck[1:]
  return c
}

func getScore(hand []card) int {
  score := 0
  numAces := 0
  for _, c := range hand {
    switch c.rank {
    case "Ace":
      score += 11
      numAces++
    case "Two":
      score += 2
    case "Three":
      score += 3
    case "Four":
      score += 4
    case "Five":
      score += 5
    case "Six":
      score += 6
    case "Seven":
      score += 7
    case "Eight":
      score += 8
    case "Nine":
      score += 9
    case "Ten", "Jack", "Queen", "King":
      score += 10
    }
  }
  for numAces > 0 && score > maxScore {
    score -= 10
    numAces --
  }
  return score
}

func main() {
  var playerHand []card
  var dealerHand []card

  // create channel
  inputCh := make(chan string)

  // Initialize the deck and shuffle it
  initDeck()
  shuffleDeck()

  // Deal initial cards
  playerHand = append(playerHand, drawCard())
  dealerHand = append(dealerHand, drawCard())
  playerHand = append(playerHand, drawCard())
  dealerHand = append(dealerHand, drawCard())

  fmt.Println("Player hand:", playerHand)
  fmt.Println("Player score:", getScore(playerHand))
  fmt.Println("Dealer hand:", dealerHand[0])

  // goroutine to take input from user
  go func() {
    for {
      var input string
      fmt.Scanln(&input)
      inputCh <- input
    }
  }()

  for getScore(playerHand) < maxScore {
    fmt.Print("Hit or stand? ")
    input := <-inputCh
    if input == "hit" {
      playerHand = append(playerHand, drawCard())
      fmt.Println("Player hand:", playerHand)
      fmt.Println("Player score:", getScore(playerHand))
    } else if input == "stand" {
      break
    } else {
      fmt.Println("Invalid input. Please enter 'hit' or 'stand'.")
    }
  }

  if getScore(playerHand) > maxScore {
    fmt.Println("Player busts. Dealer wins.")
  } else {
    fmt.Println("Dealer hand:", dealerHand)
    fmt.Println("Dealer score:", getScore(dealerHand))
    for getScore(dealerHand) < 17 {
      dealerHand = append(dealerHand, drawCard())
      fmt.Println("Dealer hand:", dealerHand)
      fmt.Println("Dealer score:", getScore(dealerHand))
    }
    if getScore(dealerHand) > maxScore {
      fmt.Println("Dealer busts. Player wins.")
    } else {
      if getScore(playerHand) > getScore(dealerHand) {
        fmt.Println("Player wins.")
      } else if getScore(playerHand) < getScore(dealerHand) {
        fmt.Println("Dealer wins.")
      } else {
        fmt.Println("It's a tie.")
      }
    }
  }
}

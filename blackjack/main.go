package main

import (
	"fmt"
	"gobasicprojects/deckpkg"
	"strings"
)

// Constants for Blackjack rules and thresholds
const (
	BlackjackScore     = 21
	DealerStandScore   = 17
	ReshuffleThreshold = 15 // Reshuffle when deckpkg has <= 15 cards
)

// Hand represents a collection of cards in the player's or dealer's hand
type Hand []deckpkg.Card

// String method for Hand to display all cards
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString hides dealer's second card
func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

// Score calculates the maximum score of the hand without busting
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deckpkg.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// MinScore calculates the minimum score of the hand treating Aces as 1
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

// min helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GameState represents the state of the game
type GameState struct {
	Deck   []deckpkg.Card
	State  State
	Player Hand
	Dealer Hand
}

// State type to define the different stages of the game
type State int8

// Game states
const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

// CurrentPlayer returns the hand of the current player based on game state
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

// clone creates a copy of the game state
func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deckpkg.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

// Shuffle shuffles the deckpkg and returns a new game state with a shuffled deckpkg
func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deckpkg.New(deckpkg.Deck(3), deckpkg.Shuffle)
	return ret
}

// Deal deals initial hands to player and dealer
func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	for i := 0; i < 2; i++ {
		ret.Player = append(ret.Player, draw(&ret.Deck))
		ret.Dealer = append(ret.Dealer, draw(&ret.Deck))
	}
	ret.State = StatePlayerTurn
	return ret
}

// Hit deals one card to the current player and returns the updated game state
func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	*hand = append(*hand, draw(&ret.Deck))
	if hand.Score() > BlackjackScore {
		return Stand(ret)
	}
	return ret
}

// Stand advances the game state to the next turn
func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

// EndHand evaluates the outcome of the game
func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", ret.Player, "\nScore:", pScore)
	fmt.Println("Dealer:", ret.Dealer, "\nScore:", dScore)
	switch {
	case pScore > BlackjackScore:
		fmt.Println("You busted")
	case dScore > BlackjackScore:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

// draw draws the top card from the deckpkg
func draw(deckpkg *[]deckpkg.Card) deckpkg.Card {
	card := (*deckpkg)[0]
	*deckpkg = (*deckpkg)[1:]
	return card
}

// PlayerTurn manages the player's turn
func PlayerTurn(gs GameState) GameState {
	var input string
	for gs.State == StatePlayerTurn {
		fmt.Println("Player:", gs.Player)
		fmt.Println("Dealer:", gs.Dealer.DealerString())
		fmt.Print("What will you do? (h)it, (s)tand, (q)uit: ")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		case "q":
			fmt.Println("Exiting game.")
			return GameState{} // return empty state to exit
		default:
			fmt.Println("Invalid option:", input)
		}
	}
	return gs
}

// DealerTurn manages the dealer's turn
func DealerTurn(gs GameState) GameState {
	for gs.State == StateDealerTurn {
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}
	return gs
}

// ReshuffleDeck checks if the deckpkg needs reshuffling and reshuffles if necessary
func ReshuffleDeck(gs GameState) GameState {
	if len(gs.Deck) <= ReshuffleThreshold {
		fmt.Println("Reshuffling the deckpkg...")
		return Shuffle(gs)
	}
	return gs
}

func main() {
	var gs GameState
	gs = Shuffle(gs)

	for {
		fmt.Println("----NEW GAME----")
		gs = ReshuffleDeck(gs)
		gs = Deal(gs)

		gs = PlayerTurn(gs)
		if gs.State != StatePlayerTurn { // Continue if not quit
			gs = DealerTurn(gs)
			gs = EndHand(gs)
		} else {
			break // Exit the game loop if the player quits
		}
	}
}

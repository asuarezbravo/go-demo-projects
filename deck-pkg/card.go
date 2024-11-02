// Package deckpkg provides utilities for creating and managing a deck of cards
// with customizable options like sorting, shuffling, filtering, adding jokers, and creating multiple decks.
package deckpkg

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents the suit of a card, with four standard suits and an optional Joker suit.
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // special case for Joker cards
)

// Rank represents the rank of a card, from Ace (1) to King (13).
type Rank uint8

const (
	_ Rank = iota // Skip zero value
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	MinRank = Ace
	MaxRank = King
)

// Card represents a playing card with a suit and rank.
type Card struct {
	Suit
	Rank
}

// CardValue returns an integer representing the card's value based on its suit and rank.
// Useful for comparing and sorting cards.
func (c Card) CardValue() int {
	return int(c.Suit)*int(MaxRank) + int(c.Rank)
}

// String returns a human-readable representation of the card, e.g., "Ace of Spades".
// For Jokers, it returns just "Joker".
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates a new deck of 52 cards (excluding Jokers by default).
// Additional options can be passed to modify the deck, such as adding jokers or shuffling.
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range [...]Suit{Spade, Diamond, Club, Heart} {
		for rank := MinRank; rank <= MaxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	// Apply all options to modify the deck (e.g., add jokers, shuffle, etc.)
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// DefaultSort sorts cards in a standard order based on suit and rank.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort returns a custom sort option that sorts cards based on a user-defined comparison function.
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less defines the default comparison function for sorting, based on CardValue.
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].CardValue() < cards[j].CardValue()
	}
}

// Shuffle randomizes the order of cards in a deck. It uses a pseudo-random number generator
// seeded with the current time to ensure different results on each shuffle.
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers adds a specified number of Joker cards to the deck.
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i), // Use Rank as a unique identifier for each Joker
				Suit: Joker,
			})
		}
		return cards
	}
}

// Filter removes cards from the deck based on a provided filter function.
// Only cards that do not match the filter condition are retained.
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// Deck creates a single deck composed of multiple decks.
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

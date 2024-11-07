
# Go Deck Package

This is a Go package for creating and managing a deck of cards. The package allows for the creation of standard 52-card decks and includes functionalities for shuffling, sorting, filtering, adding jokers, and duplicating decks.

## Features

- **Standard Deck Creation**: Generates a deck of 52 playing cards, including four suits and ranks from Ace to King.
- **Custom Deck Modifications**: Options to add jokers, duplicate decks, and filter specific cards.
- **Shuffling**: Randomly shuffles the deck using a time-based seed.
- **Sorting**: Provides default and custom sorting functions based on suit and rank.
- **Flexible Options**: Modify decks through functional options, making it easy to build customized decks.

## Usage

1. Import the package and use the `New` function to create a new deck:
   ```go
   import "your/module/path/deckpkg"
   
   deck := deckpkg.New()
   ```

2. **Deck Modifications**: Use functional options to customize the deck. Here are some examples:
   - **Shuffling**: Shuffle the deck.
     ```go
     deck := deckpkg.New(deckpkg.Shuffle)
     ```
   - **Adding Jokers**: Add a specified number of jokers.
     ```go
     deck := deckpkg.New(deckpkg.Jokers(2)) // Adds two jokers
     ```
   - **Creating Multiple Decks**: Duplicate the deck.
     ```go
     deck := deckpkg.New(deckpkg.Deck(3)) // Creates a deck with three copies of each card
     ```
   - **Filtering**: Remove specific cards.
     ```go
     deck := deckpkg.New(deckpkg.Filter(func(card deckpkg.Card) bool {
         return card.Rank == deckpkg.Ace // Removes all Aces
     }))
     ```

3. **Default Sorting**:
   Sort the deck in the default order (based on suit and rank).
   ```go
   deck = deckpkg.DefaultSort(deck)
   ```

## Functions and Options

- **New**: Initializes a new deck with optional modifications.
- **Shuffle**: Randomizes the card order in the deck.
- **DefaultSort**: Sorts the deck in default order based on suit and rank.
- **Sort**: Provides custom sorting with a user-defined comparison function.
- **Jokers**: Adds jokers to the deck.
- **Filter**: Filters out cards based on a user-defined condition.
- **Deck**: Construct a single deck composed of multiple decks.

## Design Decisions

This package was designed with modularity and flexibility in mind, allowing users to easily create and modify card decks.

### 1. Functional Options for Deck Modification
   - Functions such as `Shuffle`, `Jokers`, and `Filter` return functions that accept and modify the deck. This approach enables easy chaining of operations for highly customizable deck creation.

### 2. Custom Sorting
   - The `Sort` function allows users to pass a custom sorting function, enabling a variety of sorting behaviors beyond the default sorting order.

### 3. Default Comparison with `CardValue`
   - The `CardValue` method provides a way to quantify each card's value based on suit and rank. This method aids in default sorting and allows for easy comparison in custom sort functions.

### 4. Encapsulation of Card Details
   - The `Card` struct abstracts suit and rank, making the codebase adaptable for further extensions or custom card representations.

## Code Structure

- **`New`**: Creates a deck of 52 cards, with options for customization.
- **`Shuffle`**: Randomly shuffles the cards in the deck.
- **`DefaultSort`**: Provides a default sort order for the deck.
- **`Sort`**: Allows custom sorting of the deck.
- **`Filter`**: Filters cards based on a condition.
- **`Deck`**: Construct a single deck composed of multiple decks.

## Example Execution

Here’s an example of how to create a deck with two jokers, shuffle it, and sort it:
```go
package main

import (
    "fmt"
    "your/module/path/deckpkg"
)

func main() {
    deck := deckpkg.New(deckpkg.Jokers(2), deckpkg.Shuffle)
    fmt.Println("Shuffled Deck with Jokers:")
    for _, card := range deck {
        fmt.Println(card)
    }
    deck = deckpkg.DefaultSort(deck)
    fmt.Println("\nSorted Deck:")
    for _, card := range deck {
        fmt.Println(card)
    }
}
```

## Possible Improvements

- **Extended Card Types**: Support for additional card types or custom ranks and suits.
- **Score Calculation**: Implement scoring based on card values for game-related uses.
- **Additional Filters and Sorting Options**: Provide more filtering and sorting criteria, like sorting by color or even-odd ranks.

## License

This project is open source and available under the MIT license.
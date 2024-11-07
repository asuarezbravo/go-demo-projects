
# Go Blackjack Game

This is a command-line Blackjack game written in Go. It simulates a basic Blackjack game where the player competes against a dealer, following standard Blackjack rules. The game is designed with a modular structure and follows sequential, turn-based gameplay.

## Features

- **Dealer and Player Turns**: The game handles player and dealer actions in turn-based gameplay.
- **Deck Shuffling and Reshuffling**: The deck is shuffled initially and reshuffled automatically when it reaches a set threshold.
- **Flexible Game State Management**: The application tracks game state transitions (e.g., player's turn, dealer's turn, hand over) to ensure a smooth flow.
- **User-Friendly Prompts**: Prompts guide the player to "hit," "stand," or "quit" during their turn.
- **Scoring Logic**: Calculates scores for the player and dealer, handling Blackjack scoring rules, including Ace value adjustments.

## Usage

1. **Run the Program**: Compile or run the program directly using Go:
   ```bash
   go run blackjack.go
   ```

2. **Game Flow**:
   - The player is prompted to start a round, and two cards are dealt to both the player and the dealer.
   - During the player's turn, they can choose to "hit" (draw a card) or "stand" (end their turn).
   - If the player stands or busts (goes over 21), the dealer's turn begins. The dealer draws cards based on standard Blackjack rules.
   - The game then evaluates the final scores and displays the result.

## Packages Used

- **`deck`**: Provides a way to generate and shuffle a deck of cards for Blackjack, with helper methods to handle card properties like rank and suit.
- **`fmt`**: Used for printing game output and handling user input within the command line.
- **`strings`**: Used for string manipulation in functions that format the cards for display.

## Design Decisions

The application includes design decisions aimed at improving modularity, readability, and maintainability:

### 1. State Management with Custom Types
   - The `GameState` and `State` types manage game progression and turn tracking. This structure allows the program to manage each phase of the game, ensuring the flow from the player's turn to the dealer's turn and ultimately to score evaluation.

### 2. Deck Shuffling and Reshuffling Logic
   - The deck is shuffled at the beginning and reshuffled when there are 15 or fewer cards remaining. This ensures there are always enough cards to deal multiple rounds, similar to real-life casino practices.

### 3. Modular Turn Management
   - `PlayerTurn` and `DealerTurn` functions encapsulate player and dealer logic, improving readability and allowing easier modifications in the future.
   - `Stand` and `Hit` functions allow the player to control their turn dynamically.

### 4. Error Handling and User Prompts
   - The program provides clear prompts to guide the player and handles invalid inputs gracefully.
   - An option to quit (`q`) is included, allowing users to exit the game mid-round.

### 5. Reshuffle Function for Deck Management
   - The `ReshuffleDeck` function ensures the deck is reshuffled when it falls below a threshold. This simulates real-world deck management and prevents running out of cards during extended sessions.

## Code Structure

- **`main`**: Initializes the game and controls the flow between different phases, including shuffling, dealing, player turns, dealer turns, and score evaluation.
- **`PlayerTurn`**: Manages the player's decision to "hit" or "stand" and validates input.
- **`DealerTurn`**: Handles the dealer's actions based on Blackjack rules, including the soft 17 rule (where the dealer stands on a score of 17 with an Ace).
- **`ReshuffleDeck`**: Checks the deck's size and reshuffles if necessary.
- **`clone`**: Creates a copy of the game state, used to avoid modifying the original state directly, ensuring safe state management.

## Example Execution

```bash
go run blackjack.go
```

After starting the program, the player is prompted to either hit, stand, or quit. A round concludes when the player or dealer busts or when both have chosen to stand. The program displays the final score and the result (win, lose, or draw).

## Possible Improvements

- **Betting System**: Implementing a betting feature would add depth to the game, allowing players to wager on each round.
- **Multiple Rounds Tracking**: Track win/loss statistics across multiple rounds for longer gameplay sessions.
- **Enhanced Scoring Feedback**: Display details of wins and losses for each round to help the player understand game outcomes.
- **Hints for Players**: Offer basic hints for new players, suggesting when to hit or stand based on the player's and dealer's cards.

## License

This project is open source and available under the MIT license.
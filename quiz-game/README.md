
# Go Quiz Program

This is a command-line quiz program written in Go. It reads questions and answers from a CSV file, presents each question to the user, and evaluates responses based on their correctness within a specified time limit.

## Features

- **CSV Input**: Questions and answers are read from a CSV file specified by the user.
- **Timed Quiz**: The user can set a time limit using a flag.
- **Concurrency with Goroutines and Channels**: The program efficiently handles data input within the time limit using goroutines and channels.
- **Real-Time Scoring**: Displays the user's score upon quiz completion or when time runs out.

## Usage

1. Compile or run the program directly using Go:
   ```bash
   go run quiz.go -csv=yourfile.csv -limit=30
   ```
   Replace `yourfile.csv` with the path to your CSV file, and `30` with the desired time limit in seconds.

2. **Command-Line Flags**:
   - `-csv` (optional): Path to the CSV file containing questions and answers (default is `problems.csv`).
   - `-limit` (optional): Time limit for the quiz in seconds (default is `30` seconds).

3. **CSV File Format**: The CSV file should have questions in the first column and answers in the second, for example:
   ```csv
   question,answer
   5+5,10
   7-3,4
   ```

## Packages Used

- **`flag`**: Allows easy and effective handling of command-line arguments, making it simple to control the input CSV file and time limit.
- **`csv`**: Provides a simple way to read CSV files and structure them into a list of questions and answers using methods like `ReadAll`, which simplifies structured data conversion.
- **`time`**: Allows setting a timer to establish the time limit and manage user responses within that time using `NewTimer` and `Timer.C`.

## Design Decisions

The program incorporates various design decisions to make it modular, efficient, and user-friendly:

### 1. File Handling with `defer file.Close()`
   - Using `defer` ensures that the CSV file closes after reading, freeing up system resources and preventing memory leaks. Placing `defer` immediately after opening the file is a common Go practice for safe resource handling.

### 2. Concurrency with Goroutines and Channels
   - A goroutine is used to handle user input concurrently while the timer runs. This allows the quiz to end immediately when the time limit is reached without waiting for the user’s response.
   - A channel (`answerCh`) allows the goroutine to send the user’s response to the main process. The program selects between receiving the response from the channel or the timer signal (`timer.C`), efficiently controlling the quiz flow and response within the time limit.

### 3. Modularity and the `askQuestions` Function
   - The question and answer logic is encapsulated in an `askQuestions` function, improving code organization and facilitating future modifications, like adding new question rules or scoring configurations. `main` now focuses solely on initial setup, file reading, and calling main functions.

### 4. Data Validation and Sanitization
   - In the `parseLines` function, each answer is processed with `strings.TrimSpace` to ensure there are no unnecessary spaces that could affect answer comparison.

### 5. `exit` Function for Error Handling
   - The `exit` function centralizes error messages and program exit, allowing for a clean and consistent program termination in case of an error.

## Code Structure

- **`main`**: Initializes flags, reads the CSV file, and launches the quiz.
- **`askQuestions`**: Main function that presents each question to the user and handles the response within the time limit.
- **`parseLines`**: Converts CSV file lines into a `problem` structure, simplifying question and answer handling.
- **`exit`**: Prints an error message and terminates the program to ensure safe and controlled shutdown in case of errors.

## Example Execution

```bash
go run quiz.go -csv=problems.csv -limit=20
```

This starts a quiz using `problems.csv` with a 20-second time limit. The program will display each question and ask the user for a response. When time runs out or all questions are answered, the program shows the final score.

## Possible Improvements

- **Hints or Multiple Attempts**: To enhance educational interactivity, hints could be added, or multiple attempts could be allowed.
- **Dynamic Question Loading**: Integrate an API or database to load questions dynamically.
- **Score Breakdown**: Showing correct and incorrect answers at the end would help the user learn from their mistakes.

## License

This project is open source and available under the MIT license.

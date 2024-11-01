# Go Projects Repository

This repository contains a collection of example projects and programs developed in the Go programming language. Each project has its own directory, detailed README, and modular code structure to facilitate understanding and learning specific Go concepts.

## Projects

1. [Quiz Game](./quiz-game/README.md)
   - The **Quiz Game** is a command-line program that performs a quiz using questions and answers read from a CSV file. The program includes a configurable timer and shows a score at the end of the quiz. This project is ideal for practicing control flow, concurrency with goroutines, and file handling in Go.


2. [URL Shortener](./url-shortener/README.md)
   - The **URL Shortener** is a program that redirects URLs based on mappings defined in both hardcoded paths and a YAML configuration. The program initializes an HTTP server, allowing users to configure redirects easily and includes a fallback multiplexer for unmapped paths. This project is ideal for practicing HTTP server handling, routing, and YAML configuration parsing in Go.

## Installation and Running Projects

Each project is independent and can be run separately. To run a specific project, navigate to its directory and install any necessary dependencies before running the program. Here is an example for the **Quiz Program**:

```bash
# Clone the repository
git clone https://github.com/user/go-projects.git
cd go-projects/quiz-program

# Run the program
go run main.go -csv=problems.csv -limit=30
```

## License

This repository is licensed under the [MIT License](./LICENSE).

## Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go Examples on GitHub](https://github.com/topics/go)
- [Go Community](https://golang.org/help/)

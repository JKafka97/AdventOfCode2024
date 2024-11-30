# Advent of Code 2024 - Go Solutions

Welcome to my [Advent of Code 2024](https://adventofcode.com/2024) solutions written in **Go**!

This repository contains my solutions to each day's challenge during the Advent of Code event for the year 2024. Each day's problem is solved in its own directory with Go code, and I make use of automated tests to ensure the correctness of the solutions.

## Project Structure

The repository is structured as follows:

```
advent-of-code/
├── .github/
│   └── workflows/
│       └── go.yml              # GitHub Actions workflow for running tests
├── 2024/
│   ├── day01/
│   │   ├── main.go             # Code for Day 1 challenge
│   │   └── main_test.go        # Tests for Day 1
│   ├── day02/
│   │   ├── main.go             # Code for Day 2 challenge
│   │   └── main_test.go        # Tests for Day 2
│   └── ...                     # Other days' challenges
├── go.mod                      # Go module file for dependency management
└── go.sum                      # Go sum file for dependency integrity
```

Each day's solution is contained in a separate folder (e.g., `day01`, `day02`) with corresponding Go files for the challenge code and tests.

## Running Tests

This project uses Go's built-in testing framework to validate each solution.

To run all tests locally, simply execute:
```bash
go test ./...
```

You can also run tests for a specific day like so:
```bash
go test ./2024/day01
```

## GitHub Actions

I have set up a GitHub Actions workflow to automatically run the tests whenever changes are pushed to the repository. This ensures that my solutions are always up to date and functioning correctly.

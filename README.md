# Monkey Interpreter

This project is based on the book
[Writing An Interpreter In Go](https://interpreterbook.com/).

The most notable differences between the original book's implementation and this
one are the following:

    - Concurrency
    - More emphasis on test-driven development
    - Extending the language capabilities

## Building the Project

The project can be built locally via:

    go build -v ./...

## Running the Unit Tests

The project unit tests can be executed via:

    go test -v ./...

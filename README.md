# Monkey Interpreter

This project is based on the book
[Writing An Interpreter In Go](https://interpreterbook.com/).

The most notable differences between the original book's implementation and this
one are the following:

    - Concurrency
    - More emphasis on test-driven development
    - Extending the language capabilities

## Building the project

In a slight divergence from the idiomatic package organization used in Go, this
project's packages are contained within this repository. Therefore, the `GOPATH`
environment variable must be set to the directory where this repo was cloned.

Even better is to leverage the included `.envrc` file by installing `direnv`.

## Configuring `direnv`

The instructions for installing and using `direnv` are provided for Arch Linux.
Other GNU/Linux distributions will follow a similar pattern:

### Installation

    sudo pacman -S direnv

### Shell Configuration (zsh)

    echo 'eval "$(direnv hook zsh)"' >> ~/.zshrc

NOTE: The first time you enter the directory, you maybe prompted to allow
`direnv` to run.

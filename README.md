# JokerDB🃏: A Redis Implementaion in Golang

JokerDB🃏 is a Redis clone implemented in Golang. This project aims to provide a simplified version of Redis functionalities using the Go programming language.

## Features

- **Key-Value Store:** A basic key-value store inspired by Redis.
- **In-Memory Database:** Data is stored in-memory for simplicity.
- **Support for Data Types:** String, List, Set, etc. (as per your implementation).
- **Command-Line Interface (CLI):** Interact with JokerDB🃏 through a command-line interface.
- **Concurrency:** Implement concurrent operations for scalability.

## Getting Started

### Prerequisites

- Go (version 1.20.3)
<!-- - [Other dependencies, if any] -->

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/0x4E43/jokerdb.git
    cd jokerdb
    ```

2. Build the project:

    ```bash
    go build
    ```

3. Run JokerDB🃏:

    ```bash
    ./jokerdb
    ```

## Usage

How to use JokerDB🃏. Include examples of commands, interactions, and any important usage details.

```bash
# Example commands
jokerdb SET mykey "Hello"
jokerdb GET mykey

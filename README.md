# .cw/.cw/README.md

# .cw Programming Language

.cw is a full-stack programming language project implemented in Go, focusing on lexing, parsing, and Abstract Syntax Tree (AST) generation for easy and fast performance.

## Project Structure

```
.cw
├── cmd
│   └── main.go          # Entry point of the .cw programming language
├── internal
│   ├── lexer
│   │   ├── lexer.go     # Lexer implementation
│   │   └── lexer_test.go # Unit tests for the lexer
│   ├── parser
│   │   ├── parser.go     # Parser implementation
│   │   └── parser_test.go # Unit tests for the parser
│   ├── ast
│   │   ├── ast.go        # AST structures and methods
│   │   └── ast_test.go   # Unit tests for the AST
├── go.mod                # Go module definition
├── go.sum                # Checksums for module dependencies
└── README.md             # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd .cw
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Build the project:
   ```
   go build -o cw ./cmd/main.go
   ```

## Usage

To run the .cw programming language, use the following command:
```
./cw <source-file>
```

## Contribution Guidelines

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Open a pull request with a description of your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
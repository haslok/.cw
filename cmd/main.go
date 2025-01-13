package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "github.com/haslok/.cw/internal/lexer"
    "github.com/haslok/.cw/internal/parser"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: .cw <source_file>")
        os.Exit(1)
    }

    sourceFile := os.Args[1]
    sourceCode, err := ioutil.ReadFile(sourceFile)
    if err != nil {
        fmt.Printf("Error reading source file: %s\n", err)
        os.Exit(1)
    }

    l := lexer.New(string(sourceCode))
    p := parser.New(l)
    program := p.ParseProgram()
    if len(p.Errors()) != 0 {
        fmt.Printf("Parser errors:\n")
        for _, e := range p.Errors() {
            fmt.Printf("\t%s\n", e)
        }
        os.Exit(1)
    }

    // Here you would typically execute the program.
    fmt.Printf("Parsed program: %+v\n", program)
}
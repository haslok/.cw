package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token // Assuming Token is a struct defined in lexer.go
	}{
		{
			input:    "var myVar = 10",
			expected: []Token{ /* expected tokens */ },
		},
		{
			input:    "if myVar > 5 { print(myVar) }",
			expected: []Token{ /* expected tokens */ },
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		lexer := NewLexer(test.input) // Assuming NewLexer is a function defined in lexer.go
		for _, expectedToken := range test.expected {
			token := lexer.NextToken() // Assuming NextToken is a method defined in lexer.go
			if token != expectedToken {
				t.Errorf("Expected token %v, got %v", expectedToken, token)
			}
		}
	}
}
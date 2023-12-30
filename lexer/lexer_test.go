package lexer

import (
	"testing"

	"github.com/halfnibble/blublang/token"
)

func TestLexerNextToken(t *testing.T) {
	test_input := `=+(){},;`
	test_output_values := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lex := New(test_input)

	for i, output_value := range test_output_values {
		next_token := lex.NextToken()
		if next_token.Type != output_value.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, output_value.expectedType, next_token.Type)
		}
		if next_token.Literal != output_value.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, output_value.expectedLiteral, next_token.Literal)
		}
	}
}

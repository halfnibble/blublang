package lexer

import "github.com/halfnibble/blublang/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		// ASCII code for "NUL"
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var next_token token.Token

	switch lex.ch {
	case '=':
		next_token = newToken(token.ASSIGN, lex.ch)
	case ';':
		next_token = newToken(token.SEMICOLON, lex.ch)
	case ',':
		next_token = newToken(token.COMMA, lex.ch)
	case '+':
		next_token = newToken(token.PLUS, lex.ch)
	case '(':
		next_token = newToken(token.LPAREN, lex.ch)
	case ')':
		next_token = newToken(token.RPAREN, lex.ch)
	case '{':
		next_token = newToken(token.LBRACE, lex.ch)
	case '}':
		next_token = newToken(token.RBRACE, lex.ch)
	case 0:
		next_token.Literal = ""
		next_token.Type = token.EOF
		// default:
		// 	if isLetter(lex.ch) {
		// 		next_token.Literal = lex.readIdentifier()
		// 		next_token.Type = token.LookupIdent(next_token.Literal)
		// 		return next_token
		// 	}
		// 	if isDigit(lex.ch) {
		// 		next_token.Type = token.INT
		// 		next_token.Literal = lex.readNumber()
		// 		return next_token
		// 	}
		// 	next_token = newToken(token.ILLEGAL, lex.ch)
	}

	lex.readChar()
	return next_token
}

func newToken(token_type token.TokenType, ch byte) token.Token {
	return token.Token{Type: token_type, Literal: string(ch)}
}

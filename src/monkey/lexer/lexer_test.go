// Unit tests for the lexer package
package lexer

import (
	"monkey/token"
	"testing"
)

func TestNewLexer(t *testing.T) {
	tests := []struct {
		input string
		lexer Lexer
	}{
		{"", Lexer{"", 0, 1, 0}},
		{" ", Lexer{" ", 0, 1, ' '}},
		{"a", Lexer{"a", 0, 1, 'a'}},
		{"ab", Lexer{"ab", 0, 1, 'a'}},
	}

	for i, test := range tests {
		l := *New(test.input)
		if test.lexer != l {
			t.Errorf("tests[%d] failed, expected=%+v, got=%+v\n", i, test.lexer, l)
		}
	}
}

func TestReadChar(t *testing.T) {
	tests := []struct {
		input string
	}{
		{""},
		{" "},
		{"a"},
		{"ab"},
	}

	for i, test := range tests {
		l := New(test.input)
		for p, rp := 0, 1; p < len(test.input); p, rp = p+1, rp+1 {
			if p != l.position {
				t.Errorf("tests[%d] failed, expected l.position=%d, got l.position=%d\n", i, p, l.position)
			}
			if rp != l.readPosition {
				t.Errorf("tests[%d] failed, expected l.readPosition=%d, got l.readPosition=%d\n", i, rp, l.readPosition)
			}
			if test.input[p] != l.ch {
				t.Errorf("tests[%d] failed, expected l.ch=%q, got l.ch=%q\n", i, test.input[p], l.ch)
			}

			l.readChar()
		}
		if l.position != len(test.input) {
			t.Errorf("tests[%d] failed, expected l.position=%d, got l.position=%d\n", i, len(test.input), l.position)
		}
		if l.readPosition != len(test.input)+1 {
			t.Errorf("tests[%d] failed, expected l.position=%d, got l.position=%d\n", i, len(test.input)+1, l.position)
		}
		if l.ch != 0 {
			t.Errorf("tests[%d] failed, expected l.ch=%d, got l.ch=%q\n", i, 0, l.ch)
		}
	}
}

func TestNextTokenOperators(t *testing.T) {
	input := `=+-*/!<>`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.BANG, "!"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EOF, "EOF"},
	}

	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.expectedType {
			t.Errorf("tests[%d] - type wrong. expected=%q, got=%q\n", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q\n", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenTerminals(t *testing.T) {
	input := `(){},;#`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.ILLEGAL, "#"},
		{token.EOF, "EOF"},
	}

	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.expectedType {
			t.Errorf("tests[%d] - type wrong. expected=%q, got=%q\n", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q\n", i, test.expectedLiteral, tok.Literal)
		}
	}
}

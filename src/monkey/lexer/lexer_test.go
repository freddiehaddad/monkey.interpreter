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

func TestNextTokenWhitespace(t *testing.T) {
	input := ",\n;\t==\r! !="

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EQ, "=="},
		{token.BANG, "!"},
		{token.NOT_EQ, "!="},
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

func TestNextTokenKeywords(t *testing.T) {
	input := "let fn return if else true false"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.FUNCTION, "fn"},
		{token.RETURN, "return"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
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

func TestNextTokenIdentifiers(t *testing.T) {
	input := `lets let fns fn returns return ifs if elses else
	          trues true falses false _abc ab_c abc_`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "lets"},
		{token.LET, "let"},
		{token.IDENT, "fns"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "returns"},
		{token.RETURN, "return"},
		{token.IDENT, "ifs"},
		{token.IF, "if"},
		{token.IDENT, "elses"},
		{token.ELSE, "else"},
		{token.IDENT, "trues"},
		{token.TRUE, "true"},
		{token.IDENT, "falses"},
		{token.FALSE, "false"},
		{token.IDENT, "_abc"},
		{token.IDENT, "ab_c"},
		{token.IDENT, "abc_"},
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

func TestNextTokenLogicalOperators(t *testing.T) {
	input := `=!===!`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.NOT_EQ, "!="},
		{token.EQ, "=="},
		{token.BANG, "!"},
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

func TestNextTokenIntegers(t *testing.T) {
	input := "0 10 100"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "0"},
		{token.INT, "10"},
		{token.INT, "100"},
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

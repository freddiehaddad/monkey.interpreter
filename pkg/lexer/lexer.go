// Monkey language lexer
package lexer

import (
	"github.com/freddiehaddad/monkey.interpreter/pkg/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	tokens       chan token.Token
}

// New creates a new Lexer. The input variable specifies the data to be processed.
func New(input string) *Lexer {
	lexer := Lexer{
		input:  input,
		tokens: make(chan token.Token, 10),
	}
	lexer.readChar()

	go func(lexer *Lexer) {
		defer close(lexer.tokens)
		for t := lexer.nextToken(); t.Type != token.EOF; t = lexer.nextToken() {
			lexer.tokens <- t
		}
		lexer.tokens <- newEofToken()
	}(&lexer)

	return &lexer
}

// Advances the Lexer's input one character if we have not reached the end of input.
// Upon reaching the end of input, `ch` is set to 0 and any additional calls to
// readChar are undefined.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Returns the next character in the input without advancing.  Returns 0 when the
// end of input is reached.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Advances lexer position past all sequential whitespace characters stopping when
// reaching a non-whitespace character or the end of file.
func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

// Checks if ch is an alpha character or the underscore returning true if ch meets
// the criteria.  False otherwise.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Returns the sequence of characters matching the `isLetter` criteria.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Returns the sequence of characters up to the next ".
func (l *Lexer) readString() string {
	position := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Checks if ch is an integer character returning true if ch meets the criteria.
// False otherwise.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Returns the sequence of characters matching the `isDigit` criteria.
func (l *Lexer) readInteger() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Returns the next token in the input.
func (l *Lexer) NextToken() token.Token {
	t, ok := <-l.tokens
	if !ok {
		return newEofToken()
	}
	return t
}

// Returns the next token in the input.
func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	l.consumeWhitespace()

	switch l.ch {
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '"':
		l.readChar()
		literal := l.readString()
		if l.ch == '"' {
			tok = token.Token{Type: token.STRING, Literal: literal}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case 0:
		tok = newEofToken()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readInteger()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Returns a token representing the end of file (EOF).
func newEofToken() token.Token {
	return token.Token{Type: token.EOF, Literal: token.EOF}
}

// Returns a token of tokenType and the literal value `ch`.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

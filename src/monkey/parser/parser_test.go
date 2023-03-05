// Unit tests for the parser
package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d\n", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 12345;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does nto contain 3 statements. got=%d\n", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		if reflect.TypeOf(stmt).String() != "*ast.ReturnStatement" {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T\n", stmt)
			continue
		}

		if stmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", stmt.TokenLiteral())
		}
	}
}

func TestLetStatementErrors(t *testing.T) {
	input := `
	let x 5;
	let 10;
	let 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(p.errors) != 3 {
		t.Fatalf("expected 3 errors, got %d errrors\n", len(p.errors))
	}

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q\n", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s\n", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s\n", name, letStmt.Name.Value)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("%d errors encountered while parsing...\n", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q\n", msg)
	}

	t.FailNow()
}

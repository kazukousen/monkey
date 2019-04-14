package lexer_test

import (
	"testing"

	"github.com/kazukousen/monkey/lexer"
	"github.com/kazukousen/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		typ token.TokenType
		lit string
	}{
		{typ: token.LET, lit: "let"},
		{typ: token.IDENT, lit: "five"},
		{typ: token.ASSIGN, lit: "="},
		{typ: token.INT, lit: "5"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.LET, lit: "let"},
		{typ: token.IDENT, lit: "ten"},
		{typ: token.ASSIGN, lit: "="},
		{typ: token.INT, lit: "10"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.LET, lit: "let"},
		{typ: token.IDENT, lit: "add"},
		{typ: token.ASSIGN, lit: "="},
		{typ: token.FUNCTION, lit: "fn"},
		{typ: token.LPAREN, lit: "("},
		{typ: token.IDENT, lit: "x"},
		{typ: token.COMMA, lit: ","},
		{typ: token.IDENT, lit: "y"},
		{typ: token.RPAREN, lit: ")"},
		{typ: token.LBRACE, lit: "{"},
		{typ: token.IDENT, lit: "x"},
		{typ: token.PLUS, lit: "+"},
		{typ: token.IDENT, lit: "y"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.RBRACE, lit: "}"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.LET, lit: "let"},
		{typ: token.IDENT, lit: "result"},
		{typ: token.ASSIGN, lit: "="},
		{typ: token.IDENT, lit: "add"},
		{typ: token.LPAREN, lit: "("},
		{typ: token.IDENT, lit: "five"},
		{typ: token.COMMA, lit: ","},
		{typ: token.IDENT, lit: "ten"},
		{typ: token.RPAREN, lit: ")"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.BANG, lit: "!"},
		{typ: token.MINUS, lit: "-"},
		{typ: token.SLASH, lit: "/"},
		{typ: token.ASTERISK, lit: "*"},
		{typ: token.INT, lit: "5"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.INT, lit: "5"},
		{typ: token.LT, lit: "<"},
		{typ: token.INT, lit: "10"},
		{typ: token.GT, lit: ">"},
		{typ: token.INT, lit: "5"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.IF, lit: "if"},
		{typ: token.LPAREN, lit: "("},
		{typ: token.INT, lit: "5"},
		{typ: token.LT, lit: "<"},
		{typ: token.INT, lit: "10"},
		{typ: token.RPAREN, lit: ")"},
		{typ: token.LBRACE, lit: "{"},
		{typ: token.RETURN, lit: "return"},
		{typ: token.TRUE, lit: "true"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.RBRACE, lit: "}"},
		{typ: token.ELSE, lit: "else"},
		{typ: token.LBRACE, lit: "{"},
		{typ: token.RETURN, lit: "return"},
		{typ: token.FALSE, lit: "false"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.RBRACE, lit: "}"},
		{typ: token.INT, lit: "10"},
		{typ: token.EQ, lit: "=="},
		{typ: token.INT, lit: "10"},
		{typ: token.SEMICOLON, lit: ";"},
		{typ: token.INT, lit: "10"},
		{typ: token.NOT_EQ, lit: "!="},
		{typ: token.INT, lit: "9"},
		{typ: token.SEMICOLON, lit: ";"},
	}

	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if g, w := tok.Type, tt.typ; g != w {
			t.Errorf("%d. tokenType wrong. got %q, but want %q", i, g, w)
		}
		if g, w := tok.Literal, tt.lit; g != w {
			t.Errorf("%d. tokenLiteral wrong. got %q, but want %q", i, g, w)
		}
	}
}

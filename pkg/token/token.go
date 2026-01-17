package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	DOUBLE = "DOUBLE" // 123.44
	STRING = "STRING" // "abcd"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MULTI  = "*"
	DIV    = "/"
	REMIND = "%"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	SET      = "SET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

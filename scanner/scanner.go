package scanner

import (
	token "Lox/Token"
	"fmt"
	"strings"
)

type Scanner struct {
	Tokens  []*token.Token
	Line    int
	Start   int
	Current int
	Source  string
}

func CreateScanner(source string) *Scanner {
	return &Scanner{
		Line:   1,
		Source: source,
	}
}

func (s *Scanner) CreateTokens() ([]*token.Token, error) {
	for s.isNotAtTheEnd() {
		char := s.moveCursor()
		switch char {
		case '(':
			s.AddToken(token.LEFT_BRACKET, string(char))
		case ')':
			s.AddToken(token.RIGHT_BRACKET, string(char))
		case '{':
			s.AddToken(token.LEFT_CURLY, string(char))
		case '}':
			s.AddToken(token.RIGHT_CURLY, string(char))
		case '*':
			s.AddToken(token.MULTiPLY, string(char))
		case ';':
			s.AddToken(token.EOL, string(char))
		case '+':
			s.AddToken(token.ADDITION, string(char))
		case '-':
			s.AddToken(token.SUBRACT, string(char))
		case '/':
			s.AddToken(token.DIVIDE, string(char))
		case '%':
			s.AddToken(token.MODULUS, string(char))
		case '>':
			s.FindDoubleSymbols(">=", token.GREATER_THAN_OR_EQUAL, token.GREATER_THAN)
		case '<':
			s.FindDoubleSymbols("<=", token.LESS_THAN_OR_EQUAL, token.LESS_THAN)
		case '=':
			s.FindDoubleSymbols("==", token.DOUBLE_EQUAL, token.EQUAL)
		case '!':
			s.FindDoubleSymbols("!=", token.NOT_EQUAL, token.NEGATION)
		case '&':
			if s.Peak("&") {
				s.AddToken(token.AND, "&&")
			}
		case '|':
			if s.Peak("|") {
				s.AddToken(token.OR, "||")
			}
		case ' ':
		case '\n':
			s.Line++
		case '\t':
		case '\r':
		case '"':
			err := s.String()
			if err != nil {
				return nil, err
			}
		default:
			err := s.ProcessAlphaNumericTokens(char)
			if err != nil {
				return nil, err
			}
		}
	}
	s.Tokens = append(s.Tokens, &token.Token{
		Type: token.EOF,
		Char: "EOF",
	})

	return s.Tokens, nil
}

func (s *Scanner) ProcessAlphaNumericTokens(char rune) error {
	if s.IsAlpha(char) {
		return s.Identifier()
	}

	if s.IsDigit(char) {
		return s.Number()
	}

	return fmt.Errorf("unkonwn character [%s] at line %d, column %d, ", string(char), s.Line, s.Current)
}

func (s *Scanner) IsDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func (s *Scanner) IsAlpha(char rune) bool {
	return char >= 'a' && char <= 'z' ||
		char >= 'A' && char <= 'Z' ||
		char == '_'
}

func (s *Scanner) FindDoubleSymbols(symbols string, trueType, falseType token.Token_type) {
	if len(symbols) != 2 {
		return
	}

	if s.Peak(string(symbols[1])) {
		s.AddToken(trueType, symbols)
		return
	}

	s.AddToken(falseType, string(symbols[0]))
}

func (s *Scanner) AddToken(tokenType token.Token_type, char string) {
	s.Tokens = append(s.Tokens, &token.Token{
		Type: tokenType,
		Char: char,
	})
}

func (s *Scanner) Identifier() error {
	s.Start = s.Current - 1
	for {
		if !s.isNotAtTheEnd() {
			break
		}

		if !s.IsDigit(s.GetCurrent()) && !s.IsAlpha(s.GetCurrent()) {
			break
		}

		s.moveCursor()
	}

	return s.CheckIfWordIsReserved(s.GetSlice())
}

func (s *Scanner) CheckIfWordIsReserved(reserved string) error {
	switch reserved {
	case "if", "IF", "var", "VAR",
		"If", "Var", "Let", "Be",
		"let", "LET", "BE", "be",
		"FOR", "for", "RANGE", "range",
		"For", "Range", "continue", "Else",
		"Continue", "CONTINUE", "else", "ELSE",
		"CASE", "case", "SWITCH", "switch",
		"Case", "Switch", "String", "Integer",
		"STRING", "string", "INTEGER", "integer",
		"FLOAT", "float", "WHILE", "while",
		"Float", "While", "break", "Func",
		"Break", "BREAK", "FUNC", "func",
		"exit", "EXIT", "CLASS", "class",
		"Exit", "Class", "Super",
		"SUPER", "super", "return", "RETURN",
		"Return":
		s.AddToken(token.IDENTIFIER, strings.ToUpper(reserved))
	case "True", "TRUE", "true":
		s.AddToken(token.TRUE, "True")
	case "False", "fALSE", "false":
		s.AddToken(token.FALSE, "False")
	default:
		return fmt.Errorf("this word [%s] is not recogised ", reserved)
	}
	return nil
}

func (s *Scanner) Number() error {
	s.Start = s.Current - 1
	isFloat := false
	for {
		if !s.isNotAtTheEnd() {
			break
		}

		if !s.IsDigit(s.GetCurrent()) && s.GetCurrent() != '.' {
			break
		}

		if s.GetCurrent() == '.' {
			isFloat = true
		}

		s.moveCursor()
	}

	if isFloat {
		s.AddToken(token.FLOAT, s.GetSlice())
		return nil
	}

	s.AddToken(token.INTEGER, s.GetSlice())
	return nil
}

func (s *Scanner) String() error {
	s.Start = s.Current
	for {
		if !s.isNotAtTheEnd() {
			return fmt.Errorf("string not closed")
		}

		if s.GetCurrent() == '"' {
			break
		}

		s.moveCursor()
	}

	s.AddToken(token.STRING, s.GetSlice())
	s.Current++

	return nil
}

func (s *Scanner) GetSlice() string {
	return s.Source[s.Start:s.Current]
}

func (s *Scanner) moveCursor() rune {
	charInLine := s.GetCurrent()
	s.Current++
	return charInLine
}

func (s *Scanner) isNotAtTheEnd() bool {
	return s.Current < len(s.Source)
}

func (s *Scanner) Peak(char string) bool {
	if !s.isNotAtTheEnd() || string(s.GetCurrent()) != char {
		return false
	}
	s.moveCursor()
	return true
}

func (s *Scanner) GetCurrent() rune {
	return rune(s.Source[s.Current])
}

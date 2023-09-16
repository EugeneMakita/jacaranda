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
	End     int
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
	for s.AtTheEnd() {
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
			result, err := s.String()
			if err != nil {
				fmt.Println(err.Error())
			}
			s.AddToken(token.STRING, result)
		default:
			if !s.IsDigit(char) && !s.IsAlpha(char) {
				return nil, fmt.Errorf("unkonwn character [%s] at line %d, column %d, ", string(char), s.Line, s.Current)
			}

			if s.IsAlpha(char) {
				err := s.Identifier()
				if err != nil {
					return nil, err
				}
			}

			if s.IsDigit(char) {
				err := s.Number()
				if err != nil {
					return nil, err
				}
			}

		}
	}
	s.Tokens = append(s.Tokens, &token.Token{
		Type: token.EOF,
		Char: "EOF",
	})

	return s.Tokens, nil
}

func (s *Scanner) IsDigit(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	}

	return false
}

func (s *Scanner) IsAlpha(char rune) bool {
	if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' {
		return true
	}

	return false
}

func (s *Scanner) FindDoubleSymbols(symbols string, trueType, falseType token.Token_type) {
	if len(symbols) != 2 {
		return
	}

	if s.Peak(string(symbols[1])) {
		s.AddToken(trueType, symbols)
	} else {
		s.AddToken(falseType, string(symbols[0]))
	}
}

func (s *Scanner) AddToken(tokenType token.Token_type, char string) {
	s.Tokens = append(s.Tokens, &token.Token{
		Type: tokenType,
		Char: strings.ToUpper(string(char)),
	})
}

func (s *Scanner) Identifier() error {
	s.Start = s.Current - 1
	for {
		if s.Current >= len(s.Source) {
			break
		}

		if !s.IsDigit(rune(s.Source[s.Current])) && !s.IsAlpha(rune(s.Source[s.Current])) {
			break
		}

		s.Current++
	}

	result := s.Source[s.Start:s.Current]
	err := s.CheckIfWordIsReserved(result)
	return err
}

func (s *Scanner) CheckIfWordIsReserved(reserved string) error {
	switch reserved {
	case "if":
		s.AddToken(token.IDENTIFIER, reserved)
	case "IF":
		s.AddToken(token.IDENTIFIER, reserved)
	case "var":
		s.AddToken(token.IDENTIFIER, reserved)
	case "VAR":
		s.AddToken(token.IDENTIFIER, reserved)
	case "let":
		s.AddToken(token.IDENTIFIER, reserved)
	case "LET":
		s.AddToken(token.IDENTIFIER, reserved)
	case "BE":
		s.AddToken(token.IDENTIFIER, reserved)
	case "be":
		s.AddToken(token.IDENTIFIER, reserved)
	case "FOR":
		s.AddToken(token.IDENTIFIER, reserved)
	case "for":
		s.AddToken(token.IDENTIFIER, reserved)
	case "RANGE":
		s.AddToken(token.IDENTIFIER, reserved)
	case "range":
		s.AddToken(token.IDENTIFIER, reserved)
	case "Continue":
		s.AddToken(token.IDENTIFIER, reserved)
	case "CONTINUE":
		s.AddToken(token.IDENTIFIER, reserved)
	case "Else":
		s.AddToken(token.IDENTIFIER, reserved)
	case "ELSE":
		s.AddToken(token.IDENTIFIER, reserved)
	case "CASE":
		s.AddToken(token.IDENTIFIER, reserved)
	case "case":
		s.AddToken(token.IDENTIFIER, reserved)
	case "SWITCH":
		s.AddToken(token.IDENTIFIER, reserved)
	case "switch":
		s.AddToken(token.IDENTIFIER, reserved)
	case "STRING":
		s.AddToken(token.IDENTIFIER, reserved)
	case "string":
		s.AddToken(token.IDENTIFIER, reserved)
	case "INTEGER":
		s.AddToken(token.IDENTIFIER, reserved)
	case "integer":
		s.AddToken(token.IDENTIFIER, reserved)
	case "FLOAT":
		s.AddToken(token.IDENTIFIER, reserved)
	case "float":
		s.AddToken(token.IDENTIFIER, reserved)
	case "WHILE":
		s.AddToken(token.IDENTIFIER, reserved)
	case "while":
		s.AddToken(token.IDENTIFIER, reserved)
	case "Break":
		s.AddToken(token.IDENTIFIER, reserved)
	case "BREAK":
		s.AddToken(token.IDENTIFIER, reserved)
	case "FUNC":
		s.AddToken(token.IDENTIFIER, reserved)
	case "func":
		s.AddToken(token.IDENTIFIER, reserved)
	case "exit":
		s.AddToken(token.IDENTIFIER, reserved)
	case "EXIT":
		s.AddToken(token.IDENTIFIER, reserved)
	default:
		return fmt.Errorf("this word [%s] is not recogised ", reserved)
	}
	return nil
}

func (s *Scanner) Number() error {
	s.Start = s.Current - 1
	isFloat := false
	for {
		if s.Current >= len(s.Source) {
			break
		}

		if !s.IsDigit(rune(s.Source[s.Current])) && s.Source[s.Current] != '.' {
			break
		}

		if s.Source[s.Current] == '.' {
			isFloat = true
		}

		s.Current++
	}

	result := s.Source[s.Start:s.Current]

	if isFloat {
		s.AddToken(token.FLOAT, result)
		return nil
	}

	s.AddToken(token.INTEGER, result)
	return nil
}

func (s *Scanner) String() (string, error) {
	s.Start = s.Current
	for {
		if s.Current >= len(s.Source) {
			return "", fmt.Errorf("string not closed")
		}

		if s.Source[s.Current] == '"' {
			s.End = s.Current
			break
		}

		s.Current++
	}

	result := s.Source[s.Start:s.Current]
	s.Current++

	return result, nil
}

func (s *Scanner) moveCursor() rune {
	charInLine := s.Source[s.Current]
	s.Current++
	return rune(charInLine)
}

func (s *Scanner) AtTheEnd() bool {
	return s.Current < len(s.Source)
}

func (s *Scanner) Peak(char string) bool {
	if s.Current >= len(s.Source) || string(s.Source[s.Current]) != char {
		return false
	}

	s.Current++
	return true
}

package lexer

import (
	"bytes"
	"strings"

	"github.com/sdjdd/calc/token"
)

type DFAState int

type Lexer struct {
	tokenText strings.Builder
	tokens    []token.Token
	token     token.Token
}

func (l *Lexer) initToken(ch rune) (newState DFAState) {
	if l.tokenText.Len() > 0 {
		l.token.Text = l.tokenText.String()
		l.tokens = append(l.tokens, l.token)

		l.tokenText.Reset()
		l.token = token.Token{}
	}

	switch {
	case isAlpha(ch):
		if ch == 'i' {
			newState = Id_int1
		} else {
			newState = Id
		}
		l.token.Type = token.Identifier
		l.tokenText.WriteRune(ch)
	case isDigit(ch):
		newState = IntLiteral
		l.token.Type = token.IntLiteral
		l.tokenText.WriteRune(ch)
	case ch == '>':
		newState = GT
		l.token.Type = token.GT
		l.tokenText.WriteRune(ch)
	case ch == '+':
		newState = Plus
		l.token.Type = token.GT
		l.tokenText.WriteRune(ch)
	case ch == '-':
		newState = Minus
		l.token.Type = token.Minus
		l.tokenText.WriteRune(ch)
	case ch == '*':
		newState = Star
		l.token.Type = token.Star
		l.tokenText.WriteRune(ch)
	case ch == '/':
		newState = Slash
		l.token.Type = token.Slash
		l.tokenText.WriteRune(ch)
	case ch == ';':
		newState = Semicolon
		l.token.Type = token.Semicolon
		l.tokenText.WriteRune(ch)
	case ch == '(':
		newState = LeftParen
		l.token.Type = token.LeftParen
		l.tokenText.WriteRune(ch)
	case ch == ')':
		newState = RightParen
		l.token.Type = token.RightParen
		l.tokenText.WriteRune(ch)
	case ch == '=':
		newState = Assignment
		l.token.Type = token.Assignment
		l.tokenText.WriteRune(ch)
	default:
		newState = Initial // skip all unknown patterns
	}

	return
}

func (l *Lexer) tokenize(code string) *token.Reader {
	reader := new(bytes.Buffer)
	reader.WriteString(code)

	l.tokens = l.tokens[:0]
	l.tokenText.Reset()
	l.token = token.Token{}

	var ch rune
	state := Initial

	for reader.Len() > 0 {
		ch, _, _ = reader.ReadRune()
		switch state {
		case Initial:
			state = l.initToken(ch)
		case Id:
			if isAlpha(ch) || isDigit(ch) {
				l.tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case GT:
			if ch == '=' {
				l.token.Type = token.GE
				state = GE
				l.tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case GE, Assignment, Plus, Minus, Star, Slash,
			Semicolon, LeftParen, RightParen:
			state = l.initToken(ch)
		case IntLiteral:
			if isDigit(ch) {
				l.tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case Id_int1:
			if ch == 'n' {
				state = Id_int2
				l.tokenText.WriteRune(ch)
			} else if isAlpha(ch) || isDigit(ch) {
				state = Id
				l.tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}

		case Id_int2:
			if ch == 't' {
				state = Id_int3
				l.tokenText.WriteRune(ch)
			} else if isAlpha(ch) || isDigit(ch) {
				state = Id
				l.tokenText.WriteRune(ch)
			} else {
				state = l.initToken(ch)
			}
		case Id_int3:
			if isBlank(ch) {
				l.token.Type = token.Int
				state = l.initToken(ch)
			} else {
				state = Id
				l.tokenText.WriteRune(ch)
			}
		default:
			panic("unsupport token: " + l.tokenText.String())
		}
	}
	if l.tokenText.Len() > 0 {
		l.initToken(ch)
	}

	return token.NewReader(l.tokens)
}

package token

import (
	"strings"
)

type Token struct {
	Type TokenType // 类型
	Text string    // 文本值
}

type Reader struct {
	tokens []Token
	pos    int
}

func NewReader(tokens []Token) *Reader {
	return &Reader{
		tokens: tokens,
	}
}

func (r *Reader) Read() (token *Token) {
	if r.pos >= len(r.tokens) {
		return
	}
	token = &r.tokens[r.pos]
	r.pos++
	return
}

func (r *Reader) Peek() *Token {
	if r.pos >= len(r.tokens) {
		return nil
	}
	return &r.tokens[r.pos]
}

func (r *Reader) Unread() {
	if r.pos > 0 {
		r.pos--
	}
}

func (r *Reader) Pos() int {
	return r.pos
}

func (r *Reader) SetPos(pos int) {
	if pos >= 0 && pos < len(r.tokens) {
		r.pos = pos
	}
}

func (r *Reader) Dump() string {
	buf := strings.Builder{}
	buf.WriteString("text\t\ttype\n")
	buf.WriteString("----\t\t----\n")
	for {
		token := r.Read()
		if token == nil {
			break
		}
		buf.WriteString(token.Text + "\t\t" + token.Type.String())
		buf.WriteByte('\n')
	}

	return buf.String()
}

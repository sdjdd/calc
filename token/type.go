package token

type TokenType int

const (
	Plus  TokenType = iota // +
	Minus                  // -
	Star                   // *
	Slash                  // /

	GE // >=
	GT // >
	EQ // ==
	LE // <=
	LT // <

	Semicolon  // ;
	LeftParen  // (
	RightParen // )

	Assignment // =

	If
	Else

	Int

	Identifier // 标识符

	IntLiteral    // 整型字面量
	StringLiteral // 字符串字面量
)

func (t TokenType) String() string {
	switch t {
	case Plus:
		return "plus"
	case Minus:
		return "minus"
	case Star:
		return "star"
	case Slash:
		return "slash"
	case GE:
		return "greater than or equal"
	case GT:
		return "greater than"
	case EQ:
		return "equal"
	case LE:
		return "less than or queal"
	case LT:
		return "less than"
	case Semicolon:
		return "semicolon"
	case LeftParen:
		return "left paren"
	case RightParen:
		return "right paren"
	case Assignment:
		return "assignment"
	case If:
		return "if"
	case Else:
		return "else"
	case Int:
		return "integer"
	case Identifier:
		return "identifier"
	case IntLiteral:
		return "integer literal"
	case StringLiteral:
		return "string literal"
	default:
		return "unknown"
	}
}

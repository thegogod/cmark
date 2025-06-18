package flow

type Token = rune

const (
	Eof Token = iota

	// singles
	Comma
	Dot
	Colon
	SemiColon
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	QuestionMark

	// doubles
	DoubleColon
	ReturnType

	// arithmetic
	Plus
	PlusEq
	Minus
	MinusEq
	Star
	StarEq
	Slash
	SlashEq

	// logical
	Not
	NotEq
	Eq
	EqEq
	Gt
	GtEq
	Lt
	LtEq
	And
	Or

	// literals
	Identifier
	LString
	LByte
	LInt
	LFloat
	Nil

	// keywords
	If
	Else
	For
	Let
	Const
	Fn
	Return
	Struct
	Self
	Pub
	Use
	True
	False

	// types
	Type
	String
	Byte
	Int
	Float
	Bool
	Map
)

var Keywords = map[string]Token{
	"if":     If,
	"else":   Else,
	"for":    For,
	"let":    Let,
	"const":  Const,
	"fn":     Fn,
	"return": Return,
	"struct": Struct,
	"self":   Self,
	"pub":    Pub,
	"use":    Use,
	"true":   True,
	"false":  False,
	"string": Type,
	"byte":   Type,
	"int":    Type,
	"float":  Type,
	"bool":   Type,
	"map":    Type,
}

var Types = map[string]Token{
	"string": String,
	"byte":   Byte,
	"int":    Int,
	"float":  Float,
	"bool":   Bool,
	"map":    Map,
}

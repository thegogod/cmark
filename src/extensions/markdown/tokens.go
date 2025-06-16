package markdown

import "github.com/thegogod/cmark/tokens"

type Token = rune

const (
	Eof Token = iota

	// whitespace

	NewLine
	Space
	Tab

	// singles

	Colon             // :
	Bang              // !
	Hash              // #
	At                // @
	LeftBracket       // [
	RightBracket      // ]
	LeftParen         // (
	RightParen        // )
	LeftBrace         // {
	RightBrace        // }
	Asterisk          // *
	Plus              // +
	Percent           // %
	Dash              // -
	Underscore        // _
	Tilde             // ~
	Equals            // =
	EqualsEquals      // ==
	NotEquals         // !=
	GreaterThan       // >
	GreaterThanEquals // >=
	LessThan          // <
	LessThanEquals    // <=
	Quote             // '
	DoubleQuote       // "
	BackQuote         // `
	Period            // .
	Pipe              // |
	Or                // ||
	Ampersand         // &
	And               // &&
	Slash             // /
	BackSlash         // \

	// compounds

	Integer // 123
	Decimal // 123.4
	Text    // text
	True    // true
	False   // false
	Null    // null
)

var tokenScanners = []func(ptr *tokens.Pointer) (*tokens.Token, error){
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != 0 {
			return nil, nil
		}

		return ptr.Ok(Eof).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != ' ' {
			return nil, nil
		}

		return ptr.Ok(Space).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '\n' {
			return nil, nil
		}

		return ptr.Ok(NewLine).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '\t' {
			return nil, nil
		}

		return ptr.Ok(Tab).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != ':' {
			return nil, nil
		}

		return ptr.Ok(Colon).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '!' || ptr.Peek() != '=' {
			return nil, nil
		}

		ptr.Next()
		return ptr.Ok(NotEquals).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '!' {
			return nil, nil
		}

		return ptr.Ok(Bang).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '#' {
			return nil, nil
		}

		return ptr.Ok(Hash).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '@' {
			return nil, nil
		}

		return ptr.Ok(At).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '[' {
			return nil, nil
		}

		return ptr.Ok(LeftBracket).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != ']' {
			return nil, nil
		}

		return ptr.Ok(RightBracket).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '(' {
			return nil, nil
		}

		return ptr.Ok(LeftParen).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != ')' {
			return nil, nil
		}

		return ptr.Ok(RightParen).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '{' {
			return nil, nil
		}

		return ptr.Ok(LeftBrace).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '}' {
			return nil, nil
		}

		return ptr.Ok(RightBrace).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '*' {
			return nil, nil
		}

		return ptr.Ok(Asterisk).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '+' {
			return nil, nil
		}

		return ptr.Ok(Plus).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '%' {
			return nil, nil
		}

		return ptr.Ok(Percent).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '-' {
			return nil, nil
		}

		return ptr.Ok(Dash).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '_' {
			return nil, nil
		}

		return ptr.Ok(Underscore).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '~' {
			return nil, nil
		}

		return ptr.Ok(Tilde).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '=' || ptr.Peek() != '=' {
			return nil, nil
		}

		ptr.Next()
		return ptr.Ok(EqualsEquals).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '=' {
			return nil, nil
		}

		return ptr.Ok(Equals).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '>' || ptr.Peek() != '=' {
			return nil, nil
		}

		ptr.Next()
		return ptr.Ok(GreaterThanEquals).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '>' {
			return nil, nil
		}

		if ptr.Peek() == ' ' {
			ptr.Next()
		}

		return ptr.Ok(GreaterThan).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '<' || ptr.Peek() != '=' {
			return nil, nil
		}

		ptr.Next()
		return ptr.Ok(LessThanEquals).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '<' {
			return nil, nil
		}

		if ptr.Peek() == ' ' {
			ptr.Next()
		}

		return ptr.Ok(LessThan).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '\'' {
			return nil, nil
		}

		return ptr.Ok(Quote).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '"' {
			return nil, nil
		}

		return ptr.Ok(DoubleQuote).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '`' {
			return nil, nil
		}

		return ptr.Ok(BackQuote).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '.' {
			return nil, nil
		}

		return ptr.Ok(Period).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '|' || ptr.Peek() != '|' {
			return nil, nil
		}

		ptr.Next()
		return ptr.Ok(Or).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '|' {
			return nil, nil
		}

		return ptr.Ok(Pipe).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '&' || ptr.Peek() != '&' {
			return nil, nil
		}

		ptr.Next()
		return ptr.Ok(And).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '&' {
			return nil, nil
		}

		return ptr.Ok(Ampersand).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '/' {
			return nil, nil
		}

		return ptr.Ok(Slash).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() != '\\' {
			return nil, nil
		}

		return ptr.Ok(BackSlash).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() < '0' || ptr.Curr() > '9' {
			return nil, nil
		}

		for ptr.Peek() >= '0' && ptr.Peek() <= '9' {
			ptr.Next()
		}

		if ptr.Peek() != '.' {
			return nil, nil
		}

		if ptr.Peek() < '0' || ptr.Peek() > '9' {
			return nil, nil
		}

		for ptr.Peek() >= '0' && ptr.Peek() <= '9' {
			ptr.Next()
		}

		return ptr.Ok(Decimal).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		if ptr.Curr() < '0' || ptr.Curr() > '9' {
			return nil, nil
		}

		for ptr.Peek() >= '0' && ptr.Peek() <= '9' {
			ptr.Next()
		}

		return ptr.Ok(Integer).Ptr(), nil
	},
	func(ptr *tokens.Pointer) (*tokens.Token, error) {
		for (ptr.Peek() >= 'a' && ptr.Peek() <= 'z') || (ptr.Peek() >= 'A' && ptr.Peek() <= 'Z') {
			ptr.Next()
		}

		return ptr.Ok(Text).Ptr(), nil
	},
}

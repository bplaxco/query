package parser

import (
	"log"
	"strings"

	"github.com/bplaxco/query/pkg/exec"
)

type Lexer struct {
	appendNext bool
	data       string
	escaped    bool
	quote      rune
	tokens     []string
}

func NewLexer(data string) *Lexer {
	return &Lexer{data: data}
}

func (l *Lexer) IsQuoted() bool {
	return l.quote != 0
}

func (l *Lexer) IsQuotedWith(r rune) bool {
	return l.quote == r
}

func (l *Lexer) IsQuoteRune(r rune) bool {
	return r == '\'' || r == '"'
}

func (l *Lexer) Quote(r rune) {
	l.quote = r
}

func (l *Lexer) Unquote() {
	l.quote = 0
}

func (l *Lexer) IsEscaped() bool {
	return l.escaped
}

func (l *Lexer) Escape() {
	l.escaped = true
}

func (l *Lexer) Unescape() {
	l.escaped = false
}

func (l *Lexer) IsEscapeRune(r rune) bool {
	return r == '\\'
}

func (l *Lexer) IsWhitespaceRune(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}

func (l *Lexer) HandleEscape(r rune) bool {
	if l.IsEscapeRune(r) {
		l.Escape()
		return true
	}

	return false
}

func (l *Lexer) HandleQuote(r rune) bool {
	if l.IsQuoteRune(r) {
		if !l.IsQuoted() {
			l.Quote(r)
			return true
		}

		if l.IsQuotedWith(r) {
			l.Unquote()
			return true
		}
	}

	return false
}

func (l *Lexer) IsTokenBoundry(r rune) bool {
	return !l.IsQuoted() && l.IsWhitespaceRune(r)
}

func (l *Lexer) UpdateTokens(token string) {
	// Handle the case of an equal with spaces
	if token == "=" {
		l.tokens[len(l.tokens)-1] += token
		l.appendNext = true
	} else if l.appendNext {
		l.tokens[len(l.tokens)-1] += token
		l.appendNext = false
	} else {
		l.tokens = append(l.tokens, token)
	}
}

func (l *Lexer) Tokens() []string {
	// Guess a reasonable size for the token slice
	l.tokens = make([]string, 0, len(l.data)/5)
	var token strings.Builder

	for _, r := range l.data {
		if l.IsEscaped() {
			// Go ahead and write the rune if it's escaped
			token.WriteRune(r)
			l.Unescape()
		} else if !(l.HandleEscape(r) || l.HandleQuote(r)) {
			// There were no escapes or quote characters handled
			if l.IsTokenBoundry(r) {
				// Update the tokens since we're at a boundry
				if token.Len() > 0 {
					l.UpdateTokens(token.String())
					token.Reset()
				}
			} else {
				// Just write a normal token since nothing else
				// applies
				token.WriteRune(r)
			}
		}
	}

	return l.tokens
}

func Parse(data string) []*exec.Command {
	tokens := NewLexer(data).Tokens()
	// Guess a reasonable size for the command slice
	commands := make([]*exec.Command, 0, len(tokens)/5)

	var command *exec.Command
	for _, token := range tokens {
		if command == nil {
			command = exec.NewCommand(token)
			commands = append(commands, command)
		} else if token == "|" {
			command = nil
		} else {
			key, value, found := strings.Cut(token, "=")

			if found && len(key) > 0 {
				command.Args[key] = value
			} else {
				log.Fatalf("%s is an invalid token\n", token)
			}
		}
	}

	return commands
}

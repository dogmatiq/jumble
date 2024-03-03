package natsort

import (
	"unicode/utf8"
)

// tokenizer splits a string into a series of tokens along the boundaries of
// numeric and non-numeric characters.
type tokenizer struct {
	Input      string
	begin, end int
	tokenType  tokenType
	hasDot     bool
}

// Next returns the next token in the input string.
func (t *tokenizer) Next() token {
	for {
		if t.begin == len(t.Input) {
			return token{}
		}

		if t.end == len(t.Input) {
			return t.capture()
		}

		r, size := t.peek()

		tt := textual
		if r >= '0' && r <= '9' {
			tt = numeric
		}

		if t.tokenType != epsilon && t.tokenType != tt {
			return t.capture()
		}

		t.tokenType = tt
		t.end += size
	}
}

// capture returns the current token and resets the tokenizer state.
func (t *tokenizer) capture() token {
	tok := token{
		t: t.tokenType,
		v: t.Input[t.begin:t.end],
	}
	t.tokenType = epsilon
	t.begin = t.end
	t.hasDot = false
	return tok
}

// peek returns the next rune in the input string and its size in bytes.
func (t *tokenizer) peek() (rune, int) {
	c := t.Input[t.end]
	if c < utf8.RuneSelf {
		return rune(c), 1
	}
	return utf8.DecodeRuneInString(t.Input[t.end:])
}

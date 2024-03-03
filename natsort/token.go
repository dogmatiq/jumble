package natsort

// tokenType is an enumeration of the types of tokens that can be emitted by a
// [tokenizer].
type tokenType int

const (
	// epsilon is the "empty" token type.
	epsilon tokenType = iota

	// textual is a token type that represents a sequence of characters that can
	// not be interpreted as a number.
	textual

	// numeric is a token type that represents a sequence of characters that can
	// be interpreted as a numeric.
	numeric
)

// token is a token emitted by a [tokenizer].
type token struct {
	t tokenType
	v string
}

// compareTokens compares x to y, if the comparison is definitive, meaning that
// no more tokens need to be compared, done is true.
func compareTokens(x, y token) (cmp int, done bool) {
	if x.t == numeric && y.t == numeric {
		cmp = compareNumeric(x.v, y.v)
	} else if x.v < y.v {
		cmp = -1
	} else if x.v > y.v {
		cmp = +1
	}

	return cmp, cmp != 0 || x.t == epsilon || y.t == epsilon
}

// compareNumeric compares two numeric strings.
func compareNumeric(x, y string) int {
	ix := 0
	iy := 0

	if delta := len(x) - len(y); delta < 0 {
		ix = delta
	} else {
		iy = -delta
	}

	for ix < len(x) {
		dx := digitAt(x, ix)
		dy := digitAt(y, iy)

		if cmp := int(dx) - int(dy); cmp != 0 {
			return cmp
		}

		ix++
		iy++
	}

	return 0
}

// digitAt returns the digit at the given index in the string, or '0' if the
// index is before the beginning of the string. This simulates (non-existent)
// leading-zeroes.
func digitAt(s string, i int) byte {
	if i < 0 {
		return '0'
	}
	return s[i]
}

package natsort

// Compare returns an integer comparing two strings "naturally", such that
// numeric elements of the strings are sorted numerically, instead of
// lexigraphically.
func Compare[T ~string](x, y T) int {
	if x == y {
		return 0
	}

	tokensX := tokenizer{Input: string(x)}
	tokensY := tokenizer{Input: string(y)}

	for {
		tx := tokensX.Next()
		ty := tokensY.Next()

		if cmp, done := compareTokens(tx, ty); done {
			return cmp
		}
	}
}

// Less returns true if x sorts before y, otherwise false, such that numeric
// elements of the strings are sorted numerically, instead of lexigraphically.
func Less[T ~string](x, y T) bool {
	return Compare(x, y) < 0
}

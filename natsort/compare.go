package natsort

// Compare performs a three-way comparison of x and y. Strings of digits within
// the strings are compared using their integer value.
//
// The result will be 0 if x == y, -1 if x < y, and +1 if x > y.
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

// Less returns true if x should appear before y in a sorted list. Strings of
// digits within the strings are compared using their integer value.
func Less[T ~string](x, y T) bool {
	return Compare(x, y) < 0
}

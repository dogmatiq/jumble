package natsort_test

import (
	"testing"

	. "github.com/dogmatiq/jumble/natsort"
)

// result is used to prevent the compiler from optimizing the benchmark away.
var result bool

func BenchmarkLess(b *testing.B) {
	b.Run("empty", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = Less("", "")
		}
	})

	b.Run("text only", func(b *testing.B) {
		b.Run("equal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("alpha", "alpha")
			}
		})
		b.Run("less", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("alpha", "beta")
			}
		})
		b.Run("greater", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("beta", "alpha")
			}
		})
	})

	b.Run("numeric only", func(b *testing.B) {
		b.Run("equal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("123", "123")
			}
		})
		b.Run("less", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("123", "456")
			}
		})
		b.Run("greater", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("456", "123")
			}
		})
	})

	b.Run("alpha-numeric", func(b *testing.B) {
		b.Run("equal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("alpha 123", "alpha 123")
			}
		})
		b.Run("less", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("alpha 123", "alpha 456")
			}
		})
		b.Run("greater", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("alpha 456", "alpha 123")
			}
		})
	})

	b.Run("multi-byte characters", func(b *testing.B) {
		b.Run("equal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("αβγ 100", "αβγ 100")
			}
		})
		b.Run("less", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("αβγ 100", "αβγ 200")
			}
		})
		b.Run("greater", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = Less("αβγ 200", "αβγ 100")
			}
		})
	})
}

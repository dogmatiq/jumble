package natsort_test

import (
	"testing"

	. "github.com/dogmatiq/jumble/natsort"
)

var cases = []struct {
	Name       string
	X, Y       string
	Comparison int
}{
	{
		"text == text",
		"alpha", "alpha",
		0,
	},
	{
		"text < text",
		"alpha", "beta",
		-1,
	},
	{
		"text > text",
		"beta", "alpha",
		1,
	},

	{
		"text < text (same prefix)",
		"alpha", "alphax",
		-1,
	},
	{
		"text > text (same prefix)",
		"alphax", "alpha",
		1,
	},

	{
		"numeric == numeric",
		"1", "1",
		0,
	},
	{
		"numeric < numeric",
		"1", "2",
		-1,
	},
	{
		"numeric > numeric",
		"2", "1",
		1,
	},

	{
		"numeric == numeric (leading zeroes)",
		"001", "01",
		0,
	},
	{
		"numeric < numeric (leading zeroes)",
		"01", "2",
		-1,
	},
	{
		"numeric > numeric (leading zeroes)",
		"2", "01",
		1,
	},

	{
		"alphanumeric == alphanumeric",
		"alpha 100", "alpha 100",
		0,
	},

	{
		"alphanumeric < alphanumeric (different textual part)",
		"alpha 100", "beta 100",
		-1,
	},
	{
		"alphanumeric > alphanumeric (different textual part)",
		"beta 100", "alpha 100",
		1,
	},

	{
		"alphanumeric < alphanumeric (different numeric part)",
		"alpha 100", "alpha 200",
		-1,
	},
	{
		"alphanumeric > alphanumeric (different numeric part)",
		"alpha 200", "alpha 100",
		1,
	},

	{
		"alphanumeric == alphanumeric (multi-byte characters)",
		"αβγ 100", "αβγ 100",
		0,
	},
	{
		"alphanumeric < alphanumeric (multi-byte characters)",
		"αβγ 1", "αβγ 100",
		-1,
	},
	{
		"alphanumeric > alphanumeric (multi-byte characters)",
		"αβγ 100", "αβγ 1",
		1,
	},

	{
		"hetereogeneous < hetereogeneous",
		"alpha 100", "alpha beta",
		-1,
	},
	{
		"hetereogeneous > hetereogeneous",
		"alpha beta", "alpha 100",
		1,
	},
}

func Test_Compare(t *testing.T) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Compare(c.X, c.Y)

			if got != c.Comparison {
				t.Fatalf("unexpected result for Compare(%q, %q): got %d, want %d", c.X, c.Y, got, c.Comparison)
			}
		})
	}
}

func Test_Less(t *testing.T) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Less(c.X, c.Y)
			want := c.Comparison < 0

			if got != want {
				t.Fatalf("unexpected result for Less(%q, %q): got %t, want %t", c.X, c.Y, got, want)
			}
		})
	}
}

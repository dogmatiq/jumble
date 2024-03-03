package natsort_test

import (
	"fmt"
	"slices"

	"github.com/dogmatiq/jumble/natsort"
)

func ExampleCompare_sort() {
	values := []string{
		"version 130",
		"version 1 beta",
		"version 25",
		"version 1 alpha",
	}

	slices.SortFunc(
		values,
		natsort.Compare[string],
	)

	for _, v := range values {
		fmt.Println(v)
	}

	// Output:
	// version 1 alpha
	// version 1 beta
	// version 25
	// version 130
}

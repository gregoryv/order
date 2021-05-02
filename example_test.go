package order_test

import (
	"fmt"
	"sort"

	"github.com/gregoryv/order"
)

func ExampleLinesByPattern() {
	lines := []string{
		"bypattern.go",
		"bypattern_test.go",
		"changelog.md",
		"cmd",
		"go.mod",
		"go.sum",
		"LICENSE",
		"README.md",
	}
	patterns := []string{
		"README",
		"LICENSE",
		`\.md`,
		`go\.[mod|sum]`,
		"cmd",
	}
	sort.Sort(order.LinesByPattern(lines, patterns))
	for _, line := range lines {
		fmt.Println(line)
	}
	// output:
	// README.md
	// LICENSE
	// changelog.md
	// go.mod
	// go.sum
	// cmd
	// bypattern.go
	// bypattern_test.go
}

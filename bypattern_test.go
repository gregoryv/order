package order

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestByPattern(t *testing.T) {
	type Case struct {
		lines    []string
		patterns []string
		exp      []string
	}

	ok := func(c Case) {
		t.Helper()
		got := make([]string, len(c.lines))
		copy(got, c.lines)
		sort.Sort(LinesByPattern(got, c.patterns))
		if !reflect.DeepEqual(got, c.exp) {
			var msg strings.Builder
			msg.WriteString(fmt.Sprintf("\n    %-18s %-18s %s", "INPUT", "EXPECTED", "GOT"))
			for i, exp := range c.exp {
				msg.WriteString(fmt.Sprintf("\n%2v. %-18s %-18s %s", i+1, c.lines[i], exp, got[i]))
			}
			t.Error(msg.String())
		}
	}
	ok(Case{
		lines:    []string{"a", "b", "a", "c", "d"},
		patterns: []string{"a", "c"},
		exp:      []string{"a", "a", "c", "b", "d"},
	})

	ok(Case{
		lines: []string{
			"cmd",
			"accessprov.go",
			"components.go",
			"enrollment.go",
			"go.mod",
			"go.sum",
			"index.go",
			"README.md",
			"references.go",
			"router.go",
			"systemd.service",
			"theme.go",
		},
		patterns: []string{
			"README",
			`go\.`,
			"cmd",
			`\.go`,
		},
		exp: []string{
			"README.md",
			"go.mod",
			"go.sum",
			"cmd",
			"accessprov.go",
			"components.go",
			"enrollment.go",
			"index.go",
			"references.go",
			"router.go",
			"theme.go",
			"systemd.service",
		},
	})
}

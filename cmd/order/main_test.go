package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	t.Run("passthrough", func(t *testing.T) {
		var buf bytes.Buffer
		run(&buf, strings.NewReader("a"), "")
		if got := buf.String(); got != "a" {
			t.Errorf("got %s", got)
		}
	})

	t.Run("-f no_such_file", func(t *testing.T) {
		var buf bytes.Buffer
		run(&buf, strings.NewReader("a"), "no_such_file")
		if got := buf.String(); got != "a\n" {
			t.Errorf("got %s", got)
		}
	})

	t.Run("stdin pass through", func(t *testing.T) {
		var buf bytes.Buffer
		run(
			&buf,
			strings.NewReader("internal\nREADME\nchangelog.md\nfile.txt\n"),
			"./testdata/order",
		)
		exp := "README\ninternal\nchangelog.md\nfile.txt\n"
		if got := buf.String(); got != exp {
			t.Errorf("got %q\nexp %q", got, exp)
		}

	})

}

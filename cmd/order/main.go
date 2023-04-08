// Command order sorts lines on stdin according to patterns in the
// order file.
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/order"
)

func main() {
	var (
		cli      = cmdline.NewBasicParser()
		filename = cli.Option("-f, --pattern-files, $ORDER_PATTERN_FILES",
			"comma separated list of pattern files",
		).String("")
	)
	u := cli.Usage()
	u.Preface("Sort lines on from stdin according to patterns in the order file.")
	cli.Parse()

	run(os.Stdout, os.Stdin, filename)
}

func run(out io.Writer, in io.Reader, filename string) {

	switch {
	case filename == "":
		// no order, just pass through
		io.Copy(out, in)
		return

	default:
		// find first valid patterns file
		files := strings.Split(os.ExpandEnv(filename), ",")
	findfile:
		for _, f := range files {
			if _, err := os.Stat(f); err == nil {
				filename = f
				break findfile
			}
		}
		patterns, err := ioutil.ReadFile(filename)
		if err != nil {
			// no order file
			io.Copy(out, in)
		}

		// read stdin as lines
		var content bytes.Buffer
		io.Copy(&content, in)
		body := bytes.TrimSpace(content.Bytes())
		lines := strings.Split(string(body), "\n")

		byPattern := order.LinesByPattern(
			lines,
			strings.Split(string(patterns), "\n"),
		)
		sort.Sort(byPattern)
		for _, line := range lines {
			fmt.Fprintln(out, line)
		}
	}
}

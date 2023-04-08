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
	"github.com/gregoryv/wolf"
)

func main() {
	run(wolf.NewOSCmd())
}

func run(cmd wolf.Command) {
	var (
		cli      = cmdline.NewParser(cmd.Args()...)
		help     = cli.Flag("-h, --help")
		filename = cli.Option("-f, --pattern-files, $ORDER_PATTERN_FILES",
			"comma separated list of pattern files",
		).String("")
	)

	switch {
	case help:
		fmt.Fprint(cmd.Stdout(),
			"Sort lines on from stdin according to patterns in the order file.\n\n",
		)
		cli.WriteUsageTo(cmd.Stdout())

	case filename == "":
		// no order, just pass through
		io.Copy(cmd.Stdout(), cmd.Stdin())
		return

	default:

		// find first valid patterns file
		files := strings.Split(os.ExpandEnv(filename), ",")
		for _, f := range files {
			if _, err := os.Stat(f); err == nil {
				filename = f
				break
			}
		}
		patterns, err := ioutil.ReadFile(filename)
		if err != nil {
			// no order file
			io.Copy(cmd.Stdout(), cmd.Stdin())

		}

		// read stdin as lines
		var content bytes.Buffer
		io.Copy(&content, cmd.Stdin())
		body := bytes.TrimSpace(content.Bytes())
		lines := strings.Split(string(body), "\n")

		byPattern := order.LinesByPattern(
			lines,
			strings.Split(string(patterns), "\n"),
		)
		sort.Sort(byPattern)
		for _, line := range lines {
			fmt.Fprintln(cmd.Stdout(), line)
		}
	}
}

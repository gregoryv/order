// Command order sorts lines on stdin according to patterns in the
// order file.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/gregoryv/order"
)

func main() {
	c := &cli{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
	flag.StringVar(&c.filename, "f", "", "order file")
	flag.Parse()
	c.run()
}

type cli struct {
	io.Writer        // output of sorted stream
	io.Reader        // stream to sort
	filename  string // order file with patterns
}

func (c *cli) run() {
	if c.filename == "" {
		io.Copy(c.Writer, c.Reader)
		return
	}
	patterns, err := ioutil.ReadFile(c.filename)
	if err != nil {
		// no order file
		//fmt.Fprintln(os.Stderr, err)
		io.Copy(c.Writer, c.Reader)
		return
	}

	// read stdin as lines
	var content bytes.Buffer
	io.Copy(&content, c.Reader)
	body := bytes.TrimSpace(content.Bytes())
	lines := strings.Split(string(body), "\n")

	byPattern := order.LinesByPattern(
		lines,
		strings.Split(string(patterns), "\n"),
	)
	sort.Sort(byPattern)
	for _, line := range lines {
		fmt.Fprintln(c.Writer, line)
	}
}

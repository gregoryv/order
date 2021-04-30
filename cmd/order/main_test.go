package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_cli_bad_filename(t *testing.T) {
	var out strings.Builder
	c := &cli{
		Writer:   &out,
		Reader:   strings.NewReader("a\nb\n"),
		filename: "no_such_file",
	}
	c.run()
	exp := "a\nb\n"
	if out.String() != exp {
		t.Errorf("\ngot: %q\nexp: %q", out.String(), exp)
	}
}

func Test_cli_passthrough(t *testing.T) {
	var out strings.Builder
	input := "internal\nREADME\nchangelog.md\nfile.txt\n"
	c := &cli{
		Writer:   &out,
		Reader:   strings.NewReader(input),
		filename: "",
	}
	c.run()
	exp := "internal\nREADME\nchangelog.md\nfile.txt\n"
	if out.String() != exp {
		t.Errorf("\ngot: %q\nexp: %q", out.String(), exp)
	}
}

func Test_cli(t *testing.T) {
	patterns := "intern.*\n.*ADME\nchangelog.md"
	tmp, _ := ioutil.TempFile("", "order")
	tmp.WriteString(patterns)
	tmp.Close()
	defer os.RemoveAll(tmp.Name())
	exp := "internal\nREADME\nchangelog.md\nfile.txt\n"
	var out bytes.Buffer
	c := &cli{
		Writer:   &out,
		Reader:   strings.NewReader(exp),
		filename: tmp.Name(),
	}
	c.run()
	if out.String() != exp {
		t.Errorf("\npatterns: %q\n\ngot:\n%s\nexp:\n%s", patterns, out.String(), exp)
	}
}

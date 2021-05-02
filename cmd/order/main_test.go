package main

import (
	"os"
	"testing"

	"github.com/gregoryv/wolf"
)

func Test_main(t *testing.T) {
	os.Args = []string{"", "--help"}
	main()
}

func Test_run(t *testing.T) {

	t.Run("--help", func(t *testing.T) {
		cmd := wolf.NewTCmd("", "--help")
		defer cmd.Cleanup()
		run(cmd)
		if cmd.ExitCode != 0 {
			t.Error("exit code", cmd.ExitCode)
		}
	})

	t.Run("-f no_such_file", func(t *testing.T) {
		cmd := wolf.NewTCmd("", "-f", "no_such_file")
		defer cmd.Cleanup()
		run(cmd)
		if cmd.ExitCode != 0 {
			t.Error("exit code", cmd.ExitCode)
		}
	})

	t.Run("stdin pass through", func(t *testing.T) {
		cmd := wolf.NewTCmd()
		defer cmd.Cleanup()
		input := "internal\nREADME\nchangelog.md\nfile.txt\n"
		cmd.In.WriteString(input)
		run(cmd)
		if cmd.ExitCode != 0 {
			t.Error("exit code", cmd.ExitCode)
		}
		got := cmd.Out.String()
		if got != input {
			t.Errorf("got: %s\nexp: %s", got, input)
		}
	})

	t.Run("stdin ordered", func(t *testing.T) {
		cmd := wolf.NewTCmd("", "-f", "patterns")
		os.WriteFile("patterns", []byte("intern.*\n.*ADME\nchangelog.md"), 0644)

		cmd.In.WriteString("README\nchangelog.md\nfile.txt\ninternal\n")
		defer cmd.Cleanup()
		run(cmd)

		if cmd.ExitCode != 0 {
			t.Error("exit code", cmd.ExitCode)
		}
		got := cmd.Out.String()
		exp := "internal\nREADME\nchangelog.md\nfile.txt\n"
		if got != exp {
			t.Errorf("got: %s\nexp: %s", got, exp)
		}
	})

}

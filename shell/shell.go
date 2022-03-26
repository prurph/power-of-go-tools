package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	DryRun         bool
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	return &Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}

func (s *Session) Run() {
	input := bufio.NewReader(s.Stdin)
	for {
		fmt.Fprintf(s.Stdout, "🍉 ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.Stdout, "\nThanks for shelling!")
			break
		}
		cmd, err := CmdFromString(line)
		if err != nil { // Empty line
			fmt.Fprint(s.Stdout, "\n")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s", output)
	}
}

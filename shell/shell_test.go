package shell_test

import (
	"bytes"
	"io"
	"os"
	"shell"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString(t *testing.T) {
	t.Parallel()
	input := "/bin/ls -l main.go"
	cmd, err := shell.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromStringErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestNewSession(t *testing.T) {
	t.Parallel()
	stdin, stdout, stderr := os.Stdin, os.Stdout, os.Stderr
	want := shell.Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
	got := *shell.NewSession(stdin, stdout, stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	stdin := strings.NewReader("echo hello\n\n")
	stdout := &bytes.Buffer{}
	session := shell.NewSession(stdin, stdout, io.Discard)
	session.DryRun = true
	session.Run()
	want := "🍉 echo hello\n🍉 🍉 \nThanks for shelling!\n"
	got := stdout.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

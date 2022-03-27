package pipeline_test

import (
	"bytes"
	"errors"
	"pipeline"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStdout(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.FromString(want)
	buf := &bytes.Buffer{}
	p.Output = buf
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStdoutError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world\n")
	p.Error = errors.New("BOOM!")
	buf := &bytes.Buffer{}
	p.Output = buf
	p.Stdout()
	got := buf.String()
	if got != "" {
		t.Errorf("want no output from Stdout after error, but got %q", got)
	}
}

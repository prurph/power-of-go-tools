package findgo_test

import (
	"findgo"
	"os"
	"testing"
	"testing/fstest"
)

func TestFilesOnDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata/findgo")
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFilesInMemory(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                                {},
		"subfolder/subfolder1.go":                {},
		"subfolder/subfolder2.go":                {},
		"subfolder/subsubfolder/subsubfolder.go": {},
	}
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

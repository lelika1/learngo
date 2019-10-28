package fs_test

import (
	"testing"

	"github.com/lelika1/learngo/internal/fs"
)

func TestAddFile(t *testing.T) {
	fileSys := fs.NewTree()
	tests := []struct {
		path    string
		content string
		// "?" for failed one
		want bool
	}{
		{"/foo/boo/foo.txt", "fooboofootxt", true},
		{"/foo.txt", "footxt", true},
		{"/foo.txt", "footxt1", false},      // file already exist
		{"/foo/boo", "fooboo", false},       // incorrect path
		{"foo/boo.txt", "foobootxt", false}, // incorrect path
	}
	for _, test := range tests {
		got := fileSys.AddFile(test.path, test.content)
		if got != test.want {
			t.Errorf("AddFile(%v, %v) = %v, want %v", test.path, test.content, got, test.want)
		}
	}
}

func TestReadFile(t *testing.T) {
	fileSys := fs.NewTree()
	fileSys.AddFile("/foo/foo.txt", "foofootxt")
	fileSys.AddFile("/foo/foo1.txt", "foofoo1txt")
	fileSys.AddFile("/foo/boo/foo.txt", "fooboofootxt")
	fileSys.AddFile("/foo.txt", "footxt")

	tests := []struct {
		path string
		// "?" for failed one
		want string
	}{
		{"/foo/boo/foo.txt", "fooboofootxt"},
		{"/foo.txt", "footxt"},
		{"/foo/foo1.txt", "foofoo1txt"},
		{"/foo/boo/foo", "?"},
		{"/foo/foo11.txt", "?"},
		{"/", "?"},
	}
	for _, test := range tests {
		got, ok := fileSys.ReadFile(test.path)
		if !ok {
			if test.want != "?" {
				t.Errorf("ReadFile(%v) = (%v:%v), want %v", test.path, got, ok, test.want)
			}
			continue
		}

		if got != test.want {
			t.Errorf("ReadFile(%v) = (%v:%v), want %v", test.path, got, ok, test.want)
		}
	}
}

func TestListDir(t *testing.T) {
	fileSys := fs.NewTree()
	fileSys.AddFile("/foo/foo.txt", "foofootxt")
	fileSys.AddFile("/foo/bar.txt", "foofoo1txt")
	fileSys.AddFile("/foo/bar/foo.txt", "fooboofootxt")
	fileSys.AddFile("/foo.txt", "footxt")

	tests := []struct {
		dir string
		// Format: "[{"foo.txt"}, {"nested_dir", dir=true}]" for successful result; "?" for failed one.
		want string
	}{
		{"/foo/bar", "[{\"foo.txt\"}]"},
		{"/foo", "[{\"foo.txt\"}, {\"bar.txt\"}, {\"bar\", dir=true}]"},
		{"/", "[{\"foo\", dir=true}, {\"foo.txt\"}]"},
		{"/foo.txt", "?"},
		{"foo/bar/bar2", "?"},
		{"/foo/foo1", "?"},
	}
	for _, test := range tests {
		entries, ok := fileSys.ListDir(test.dir)
		got := "["
		for _, entry := range entries {
			if got != "[" {
				got += ", "
			}
			got = got + entry.String()
		}
		got += "]"
		if !ok {
			if test.want != "?" {
				t.Errorf("ListDir(%v) = (%v:%v), want %v", test.dir, got, ok, test.want)
			}
			continue
		}

		if got != test.want {
			t.Errorf("ListDir(%v) = (%v:%v), want %v", test.dir, got, ok, test.want)
		}
	}
}

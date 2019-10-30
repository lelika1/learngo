package fs_test

import (
	"fmt"
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
		{"/foo.txt", "footxt", true},
		{"/foo.txt", "footxt1", false}, // file already exist
		{"/foo/boo/foo.txt", "fooboofootxt", true},
		{"/foo/boo", "fooboo", false},       // there is folder with name boo
		{"foo/boo.txt", "foobootxt", false}, // incorrect path
		{"/foo/boo/foo", "fooboofoo", true},
		{"/foo/boo/foo/txt.txt", "fooboofootxttxt", false}, // part of path is a file
		{"", "d", false},                                   // incorrect path
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
	fileSys.AddFile("/bar.txt/bar", "bartxtbar")

	tests := []struct {
		path string
		// "?" for failed one
		want string
	}{
		{"/foo/boo/foo.txt", "fooboofootxt"},
		{"/foo.txt", "footxt"},
		{"/foo/foo1.txt", "foofoo1txt"},
		{"/bar.txt/bar", "bartxtbar"},
		{"/foo/boo/foo", "?"},   // file doesn't exist
		{"/foo/foo11.txt", "?"}, // file doesn't exist
		{"/foo", "?"},           // read folder not file
		{"/", "?"},              // read folder not file
		{"/bar.txt", "?"},       // read folder not file
		{"/bar.txt/ bar", "?"},  //  incorrect path
		{"bar.txt/bar", "?"},    //  incorrect path
		{"", "?"},               //  incorrect path
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
	fileSys.AddFile("/foo.txt", "footxt")
	fileSys.AddFile("/foo/foo.txt", "foofootxt")
	fileSys.AddFile("/foo/aaa.txt", "fooaaatxt")
	fileSys.AddFile("/foo/boo/foo.txt", "fooboofootxt")
	fileSys.AddFile("/foo/bar.txt", "foofoo1txt")
	fileSys.AddFile("/foo/bar/bar", "foobarbar")
	fileSys.AddFile("/foo/bar/foo.txt", "foobarfootxt")
	fileSys.AddFile("/bar.txt/bar", "bartxtbar")

	tests := []struct {
		dir string
		// Format: "[{"foo.txt"}, {"nested_dir", dir=true}]" for successful result; "?" for failed one.
		want string
	}{
		{"/foo/bar", "[{\"bar\"}, {\"foo.txt\"}]"},
		{"/foo", "[{\"aaa.txt\"}, {\"bar\", dir=true}, {\"bar.txt\"}, {\"boo\", dir=true}, {\"foo.txt\"}]"},
		{"/", "[{\"bar.txt\", dir=true}, {\"foo\", dir=true}, {\"foo.txt\"}]"},
		{"/bar.txt///", "[{\"bar\"}]"},
		{"foo/bar/bar2", "?"}, // incorrect path
		{"/bar.txt/bar", "?"}, // it is a file, not a folder
		{"/foo.txt", "?"},     // it is a file, not a folder
		{"/foo/foo1", "?"},    // folder doesn't exist
		{"", "?"},             // incorrect path
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

func TestNormalizeString(t *testing.T) {

	tests := []struct {
		path string
		want string
	}{
		{"/foo/bar", "/foo/bar"},
		{"/foo///bar", "/foo/bar"},
		{"//foo///bar", "/foo/bar"},
		{"/foo///bar", "/foo/bar"},
		{"/foo///bar/////", "/foo/bar"},
		{"/foo///bar/////    ", "/foo/bar"},
		{"/foo/bar     ", "/foo/bar"},
		{"/////foo     ", "/foo"},
		{"///", ""},
		{"///   ", ""},
		{"//", ""},
		{"/", ""},
	}
	for _, test := range tests {
		got := fs.NormalizePath(test.path)
		if got != test.want {
			t.Errorf("NormalizePath(%v) = %v, want %v", test.path, got, test.want)
		}
	}
}

func TestString(t *testing.T) {
	fileSys := fs.NewTree()
	fileSys.AddFile("/one/new/few", "")
	fileSys.AddFile("/one/new/many", "")
	fileSys.AddFile("/one/old/few", "")
	fileSys.AddFile("/one/old/many", "")
	fileSys.AddFile("/two/new/few", "")
	fileSys.AddFile("/two/new/many", "")
	fileSys.AddFile("/two/old/few", "")
	fileSys.AddFile("/two/old/many", "")
	fmt.Println(fileSys)
}

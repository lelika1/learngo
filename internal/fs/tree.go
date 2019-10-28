package fs

import (
	"strings"
)

// Tree ...
type Tree struct {
	root *Entry
}

// NewTree ...
func NewTree() *Tree {
	var tree Tree
	tree.root = &Entry{Name: "", Directory: true}
	return &tree
}

// AddFile adds a new file with the given path and content to a file tree.
// Returns false if the file with the given path already exists.
func (t *Tree) AddFile(path, content string) (ok bool) {
	if !strings.HasSuffix(path, ".txt") || path[0] != '/' {
		return false
	}

	pathPart := strings.Split(path, "/")

	existFolder := 0
	lastCommon := t.root
	for ; existFolder < len(pathPart); existFolder++ {
		for _, child := range lastCommon.children {
			if child.Name == pathPart[existFolder] {
				lastCommon = child
				break
			}
		}
		if lastCommon.Name != pathPart[existFolder] {
			break
		}
	}
	if existFolder == len(pathPart) {
		return false // file was found
	}
	if existFolder < len(pathPart) {
		for i := existFolder; i < len(pathPart)-1; i++ {
			newFold := &Entry{
				Name:      pathPart[i],
				Directory: true,
			}
			lastCommon.children = append(lastCommon.children, newFold)
			lastCommon = newFold
		}
		newFile := &Entry{
			Name:    pathPart[len(pathPart)-1],
			content: content,
		}
		lastCommon.children = append(lastCommon.children, newFile)
	}

	return true
}

// ReadFile returns content for a given file by its path.
// Returns false if the file does not exist.
func (t *Tree) ReadFile(path string) (content string, ok bool) {
	if !strings.HasSuffix(path, ".txt") || path[0] != '/' {
		return "s", false
	}

	pathPart := strings.Split(path, "/")

	existFolder := 0
	lastCommon := t.root
	for ; existFolder < len(pathPart); existFolder++ {
		for _, child := range lastCommon.children {
			if child.Name == pathPart[existFolder] {
				lastCommon = child
				break
			}
		}
		if lastCommon.Name != pathPart[existFolder] {
			break
		}
	}
	if existFolder == len(pathPart) {
		return lastCommon.content, true // file was found
	}
	return "s", false
}

// Entry describes one directory entry (either a file or a nested directory).
type Entry struct {
	// Name of a file (i.e "test.txt") or a directory (i.e. "directory").
	Name      string
	Directory bool // if this is not a file, but a directory
	content   string
	children  []*Entry
}

func (e *Entry) String() string {
	result := "{\"" + e.Name + "\""
	if e.Directory {
		result += ", dir=true"
	}
	result += "}"
	return result
}

// ListDir returns a slice of entries that are directly inside a given directory.
// Returns false if the dir does not exist or has wrong format.
// Example: ListDir("/foo/bar") -> [{"foo.txt"}, {"bar.txt"}, {"nested_dir", dir=true}]
func (t *Tree) ListDir(dir string) (entries []*Entry, ok bool) {
	if strings.HasSuffix(dir, ".txt") || dir[0] != '/' {
		return []*Entry{}, false
	}

	pathPart := strings.Split(dir, "/")
	lastCommon := t.root
	for i := 0; i < len(pathPart); i++ {
		for _, child := range lastCommon.children {
			if child.Name == pathPart[i] {
				lastCommon = child
				break
			}
		}
		if lastCommon.Name != pathPart[i] {
			return []*Entry{}, false
		}
	}
	return lastCommon.children, true
}

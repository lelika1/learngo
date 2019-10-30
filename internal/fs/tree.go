package fs

import (
	"sort"
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
	tree.root.children = make(map[string]*Entry)
	return &tree
}

// NormalizePath delete redundant /
func NormalizePath(path string) string {
	path = strings.TrimSpace(path)
	for i := strings.Index(path, "//"); i != -1; i = strings.Index(path, "//") {
		path = path[:i] + path[i+1:]
	}

	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path
}

// AddFile adds a new file with the given path and content to a file tree.
// Returns false if the file with the given path already exists.
func (t *Tree) AddFile(path, content string) (ok bool) {
	if len(path) == 0 || path[0] != '/' {
		return false
	}
	pathPart := strings.Split(NormalizePath(path), "/")

	i := 1
	curFold := t.root
	for ; i < len(pathPart); i++ {
		child, ok := curFold.children[pathPart[i]]
		if ok {
			if !child.Directory {
				return false // where is a file in the middle of path
			}
			curFold = child
		} else {
			break
		}
	}

	if i == len(pathPart) {
		return false // file already exist in file system
	}

	for j := i; j < len(pathPart)-1; j++ {
		newFold := &Entry{
			Name:      pathPart[j],
			Directory: true,
		}
		newFold.children = make(map[string]*Entry)
		curFold.children[pathPart[j]] = newFold
		curFold = newFold
	}

	newFile := &Entry{
		Name:    pathPart[len(pathPart)-1],
		content: content,
	}
	curFold.children[pathPart[len(pathPart)-1]] = newFile

	return true
}

// ReadFile returns content for a given file by its path.
// Returns false if the file does not exist or format is incorrect.
func (t *Tree) ReadFile(path string) (content string, ok bool) {
	if len(path) == 0 || path[0] != '/' {
		return "", false
	}

	pathPart := strings.Split(NormalizePath(path), "/")
	curFold := t.root
	for i := 1; i < len(pathPart); i++ {
		if child, ok := curFold.children[pathPart[i]]; ok {
			curFold = child
		} else {
			return "", false
		}
	}
	if curFold.Directory {
		return "", false
	}
	return curFold.content, true
}

// Entry describes one directory entry (either a file or a nested directory).
type Entry struct {
	// Name of a file (i.e "test.txt") or a directory (i.e. "directory").
	Name      string
	Directory bool // if this is not a file, but a directory
	content   string
	children  map[string]*Entry
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
	if len(dir) == 0 || dir[0] != '/' {
		return nil, false
	}

	pathPart := strings.Split(NormalizePath(dir), "/")
	curFold := t.root
	for i := 1; i < len(pathPart); i++ {
		if child, ok := curFold.children[pathPart[i]]; ok {
			curFold = child
		} else {
			return nil, false
		}
	}
	if curFold != t.root && !curFold.Directory {
		return nil, false
	}

	return sorted(curFold.children), true
}

func sorted(entries map[string]*Entry) (result []*Entry) {
	for _, v := range entries {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func (t *Tree) String() string {
	var sb strings.Builder
	type folderInStack struct {
		name  string
		edges []bool
		desc  *Entry
	}
	var stack []folderInStack
	stack = append(stack, folderInStack{"root", []bool{}, t.root})
	for len(stack) != 0 {
		curFold := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i := 0; i < len(curFold.edges)-1; i++ {
			if curFold.edges[i] {
				sb.WriteString("|  ")
			} else {
				sb.WriteString("   ")
			}
		}
		if len(curFold.edges) != 0 {
			sb.WriteString("|\n")
		}

		for i := 0; i < len(curFold.edges)-1; i++ {
			if curFold.edges[i] {
				sb.WriteString("|  ")
			} else {
				sb.WriteString("   ")
			}
		}
		if len(curFold.edges) != 0 {
			if !curFold.edges[len(curFold.edges)-1] {
				sb.WriteString("\\__")
			} else {
				sb.WriteString("|__")
			}
		}
		if curFold.desc.Directory {
			sb.WriteString("📁")
		}
		sb.WriteString(curFold.name + "\n")

		if curFold.desc.Directory {
			nested := sorted(curFold.desc.children)
			for i := len(nested) - 1; i >= 0; i-- {
				edges := make([]bool, len(curFold.edges))
				copy(edges, curFold.edges)
				edges = append(edges, i != len(nested)-1)
				stack = append(stack, folderInStack{
					name:  nested[i].Name,
					edges: edges,
					desc:  nested[i],
				})
			}
		}
	}
	return sb.String()
}

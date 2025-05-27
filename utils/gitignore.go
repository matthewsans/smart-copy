package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	enry "github.com/go-enry/go-enry/v2"
	ignore "github.com/sabhiram/go-gitignore"
)

const maxSrc = 100 * 1024 // 100 KB

// Matcher decides whether a path should be ignored.
type Matcher struct {
	stack []frame
}

type frame struct {
	base  string
	ign   *ignore.GitIgnore
	depth int
}

func New() *Matcher { return &Matcher{} }

// UpdateStack updates the stack when you descend/ascend during a walk.
func (m *Matcher) UpdateStack(path string, d fs.DirEntry) error {
	curDepth := depth(path)
	// pop
	for len(m.stack) > 0 && m.stack[len(m.stack)-1].depth >= curDepth {
		m.stack = m.stack[:len(m.stack)-1]
	}
	// push

	if d.IsDir() {
		gi := filepath.Join(path, ".gitignore")
		if info, err := os.Stat(gi); err == nil && !info.IsDir() {
			if ign, err := ignore.CompileIgnoreFile(gi); err == nil {
				m.stack = append(m.stack, frame{
					base:  path,
					ign:   ign,
					depth: curDepth,
				})
			}
		}
	}
	return nil
}

// Ignored reports whether the given path matches any stacked ignores.
func (m *Matcher) Ignored(path string) bool {
	for i := len(m.stack) - 1; i >= 0; i-- {
		rel, err := filepath.Rel(m.stack[i].base, path)
		if err == nil && m.stack[i].ign.MatchesPath(rel) {
			return true
		}
	}

	return false
}

func depth(p string) int {
	return strings.Count(filepath.Clean(p), string(os.PathSeparator))
}

func isBinary(path string) bool {
	f, err := os.Open(path)
	if err != nil { // unreadable → treat as non-binary so we still walk
		return false
	}
	defer f.Close()

	buf := make([]byte, 8_000) // Linguist & enry only look at the first 8 kB
	n, _ := f.Read(buf)
	return enry.IsBinary(buf[:n])
}

func shouldKeep(path string, d fs.DirEntry) bool {
	if d.IsDir() {
		return !enry.IsVendor(path) && !enry.IsDotFile(path)
	}

	// fast reject
	if isBinary(path) || enry.IsImage(path) || strings.HasSuffix(path, ".svg") || enry.IsVendor(path) {
		return false
	}

	// keep READMEs & small docs
	if enry.IsDocumentation(path) || enry.IsConfiguration(path) {
		if info, _ := d.Info(); info.Size() <= maxSrc {
			return true
		}
	}

	// keep real code
	blob, _ := os.ReadFile(path)
	lang := enry.GetLanguage(path, blob)
	if lang == "" { // unknown
		return false
	}
	group := enry.GetLanguageGroup(lang)       // e.g. SCSS → CSS
	if group == "Markup" || group == "Prose" { // markdown, etc.
		return false
	}
	if info, _ := d.Info(); info.Size() > maxSrc {
		return false // huge generated code / minified bundle
	}
	return true
}

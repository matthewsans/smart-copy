// ── copy.go ──
package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// walkAll just walks every file/dir under root, calling fn on each.
func walkAll(root string, fn func(string, fs.DirEntry) error) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		return fn(path, d)
	})
}

// walkWithGitIgnore walks like WalkDir, but skips anything the Matcher says to ignore.
func walkWithGitIgnore(root string, fn func(string, fs.DirEntry) error) error {
	m := New()
	root = filepath.Clean(root)

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if err := m.UpdateStack(path, d); err != nil {
			return err
		}

		// skip ignored paths (but never skip the root itself)
		if path != root && m.Ignored(path) || !shouldKeep(path, d) {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		return fn(path, d)
	})
}

// SmartCopy either walks everything or honors .gitignore, printing each “kept” path.
func SmartCopy(root string, useIgnore bool) (string, error) {
	// define your “do work” callback here
	var buf bytes.Buffer
	cb := func(path string, d fs.DirEntry) error {
		appendFile(&buf, path)
		fmt.Println("KEEP:", path)
		return nil
	}

	var err error
	if useIgnore {
		err = walkWithGitIgnore(root, cb)
	} else {
		err = walkAll(root, cb)
	}
	if err != nil {
		return "", err
	}
	fmt.Println(buf.String())
	return "done", nil
}

func appendFile(buf *bytes.Buffer, path string) error {
	buf.WriteString("\n \n============================================== \n \n" +
		"============================================== \n \n ")

	buf.WriteString("// ---- " + filepath.Base(path) + " ----\n \n")
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(buf, f) // io.Copy handles chunking & growth
	return err
}

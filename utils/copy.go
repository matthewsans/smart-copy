package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	clip "github.com/atotto/clipboard"
)

var currRoot string

// Walks every file/dir under root, calling fn on each.
func walkAll(root string, fn func(string, fs.DirEntry) error) error {
	currRoot = filepath.Clean(root)

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return fn(path, d)
	})
}

// Skips paths based on .gitignore and optionally shouldKeep.
func walkWithGitIgnore(root string, fn func(string, fs.DirEntry) error, useShouldKeep bool) error {
	m := New()
	root = filepath.Clean(root)
	currRoot = filepath.Clean(root)

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if err := m.UpdateStack(path, d); err != nil {
			return err
		}

		// Skip ignored paths (but never skip the root itself)
		if path != root && m.Ignored(path) {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		// Apply shouldKeep filter only if useShouldKeep is true
		if useShouldKeep && !shouldKeep(path, d) {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			return nil // Continue walking into the directory
		}
		return fn(path, d) // Call fn for files
	})
}

// SmartCopy handles copying based on the provided flags.
func SmartCopy(root string, useGitignore bool, useShouldKeep bool, listOnly bool) (string, error) {
	var buf bytes.Buffer

	cb := func(path string, d fs.DirEntry) error {
		rel, err := filepath.Rel(currRoot, path)
		if err != nil {
			rel = path
		}

		if listOnly {
			buf.WriteString(rel + "\n")
		} else {
			appendFile(&buf, path)
			fmt.Println("KEEP:", path)
		}
		return nil
	}

	var err error
	if useGitignore {
		err = walkWithGitIgnore(root, cb, useShouldKeep)
	} else {
		err = walkAll(root, cb)
	}
	if err != nil {
		return "", err
	}

	fmt.Printf("DEBUG len(buf) = %d bytes\n", buf.Len())

	if err := clip.WriteAll(buf.String()); err != nil { // 👈 capture error
		fmt.Fprintf(os.Stderr, "CLIPBOARD-ERR: %v\n", err)
		return "", err
	}

	return "done", nil
}

func appendFile(buf *bytes.Buffer, path string) error {
	rel, err := filepath.Rel(currRoot, path)
	if err != nil {
		rel = path
	}
	buf.WriteString("\n // ---- " + rel + " ----\n \n")
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(buf, f) // io.Copy handles chunking & growth
	return err
}

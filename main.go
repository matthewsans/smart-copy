package main

import (
	"fmt"
	"os"

	"github.com/mango/smart-copy/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: smartcopy <path> [--mode <mode>]")
		fmt.Println("Modes: minimal, gitignore, all")
		os.Exit(1)
	}

	var path string = os.Args[1]
	var mode string = "minimal" // Default to Minimal Mode

	// Parse --mode flag
	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] == "--mode" && i+1 < len(os.Args) {
			mode = os.Args[i+1]
			break
		}
	}

	fmt.Println(mode)
	// Set flags based on mode
	var useGitignore, useShouldKeep, listOnly bool
	switch mode {
	case "minimal":
		useGitignore = true
		useShouldKeep = true
	case "gitignore":
		useGitignore = true
		useShouldKeep = false
	case "all":
		useGitignore = false
		useShouldKeep = false
	case "gitignore-list":
		useGitignore = true
		useShouldKeep = false
		listOnly = true
	default:
		fmt.Println("Invalid mode:", mode)
		os.Exit(1)
	}

	_, err := utils.SmartCopy(path, useGitignore, useShouldKeep, listOnly)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println("✔ Copied to clipboard!")
}

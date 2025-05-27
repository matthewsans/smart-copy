package main

import (
	"fmt"
	"os"

	"github.com/mango/smart-copy/utils"
)

var useIgnore bool = false

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: smartcopy <path> [--gitignore]")
		os.Exit(1)
	}

	var path string = os.Args[1]
	useIgnore = len(os.Args) > 2 && os.Args[2] == "--gitignore"

	_, err := utils.SmartCopy(path, useIgnore)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println("✔ Copied to clipboard!")

}

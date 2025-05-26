package utils

import (
	"fmt"
	"log"
	"os"
)

// Returns true for file, false for path
func isFile(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("file")
		return false, nil
	case mode.IsRegular():
		fmt.Println("file")
		return true, nil
	}

	return false, err
}

func SmartCopy(path string, useIgnore bool) (string, error) {

	fmt.Println(path)

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
	return " ", nil

}

package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error opening file in path: %s", path)
		os.Exit(1)
	}

	return string(f)
}

func PrintArray(arr []string) {
	for i, item := range arr {
		fmt.Println(i, ">", item)
	}
}

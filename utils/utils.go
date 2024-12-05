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

func IsSameArray(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func Contains(arr []string, el string) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}

	return false
}

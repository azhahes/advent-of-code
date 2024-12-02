package file

import (
	"fmt"
	"os"
)

func Open(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %s, Error: %s", filePath, err.Error())
		panic(err)
	}
	return f
}

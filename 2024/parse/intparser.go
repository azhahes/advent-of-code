package parse

import (
	"fmt"
	"strconv"
)

func ParseInt(val string) int {
	a, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		fmt.Printf("Failed to parse int from %s", val)
		panic(err)
	}
	return int(a)
}

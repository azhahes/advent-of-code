package parse

import (
	"fmt"
	"strconv"
)

func MustToInt(val string) int {
	a, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		fmt.Printf("Failed to parse int from %s", val)
		panic(err)
	}
	return int(a)
}

func MustToIntSlice(val []string) []int {
	result := make([]int, 0, len(val))
	for _, v := range val {
		a, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Printf("Failed to parse int from %s", v)
			panic(err)
		}
		result = append(result, int(a))
	}
	return result
}

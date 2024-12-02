package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	scanner := bufio.NewScanner(f)
	left, right := make([]int, 0), make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		a, err := strconv.ParseInt(words[0], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := strconv.ParseInt(words[1], 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		left = append(left, int(a))
		right = append(right, int(b))
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i, v := range left {
		diff := int(math.Abs(float64(v - right[i])))
		sum += diff
	}
	fmt.Println(sum)
}

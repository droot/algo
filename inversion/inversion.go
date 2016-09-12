package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func countInversions(a []int) int {
	_, count := countInversionInternal(a)
	return count
}

func countInversionInternal(a []int) ([]int, int) {
	if len(a) <= 1 {
		return a, 0
	}

	mid := len(a) / 2

	first := a[:mid] // till mid index, mid is excluded
	rest := a[mid:]  // mid onwards, mid is included

	sortedFirst, invervisonsInFirst := countInversionInternal(first)
	sortedRest, invervisonsInRest := countInversionInternal(rest)

	sorted, splitInversions := mergeAndCountSplitInversions(sortedFirst, sortedRest)
	return sorted, invervisonsInFirst + invervisonsInRest + splitInversions
}

func mergeAndCountSplitInversions(first, rest []int) (merged []int, count int) {
	merged = make([]int, len(first)+len(rest))
	i := 0
	j := 0
	k := 0
	for ; k < len(first)+len(rest) && i < len(first) && j < len(rest); k++ {
		if first[i] < rest[j] {
			merged[k] = first[i]
			i++
		} else {
			// split conversion here
			count += len(first) - i
			merged[k] = rest[j]
			j++
		}
	}

	for ; i < len(first); i++ {
		merged[k] = first[i]
		k++
	}

	for ; j < len(rest); j++ {
		merged[k] = rest[j]
		k++
	}
	return
}

// readFile reads the numbers from the given file and returns array of numbers.
func readFile(fileName string) (numbers []int, err error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")

	numbers = []int{}
	for _, line := range lines {
		var n int
		line = strings.Trim(line, "\r")
		if line == "" {
			continue
		}
		n, err = strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func main() {
	numbers, err := readFile("testdata.txt")
	if err != nil {
		fmt.Printf("error reading file testdata.txt :: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Read %d numbers from file... peek at the array: %v\n",
		len(numbers), numbers[:5])

	inversions := countInversions(numbers)
	fmt.Printf("number of inversions: %d \n", inversions)
}

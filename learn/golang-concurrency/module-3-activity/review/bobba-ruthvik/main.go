package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var PARTITION_SIZE int = 4

func sortArray(partition []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting partition: ", partition)
	sort.Ints(partition)
	fmt.Println("Sorted partition: ", partition)
}

func mergePartitions(partitions [][]int) []int {
	var result []int
	for _, partition := range partitions {
		result = merge(result, partition)
	}
	return result
}

func merge(a, b []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}
	return result
}

func main() {
	fmt.Println("Enter a space separated series of integers: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	strArr := strings.Split(input, " ")
	n := len(strArr)
	intArr := make([]int, n)

	for i, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}
		intArr[i] = num
	}

	fmt.Println("Converted Integer Array: ", intArr)
	var partitionedArr [][]int

	for i := 0; i < n; i += PARTITION_SIZE {
		end := i + PARTITION_SIZE
		if end > n {
			end = n
		}
		partitionedArr = append(partitionedArr, intArr[i:end])
	}

	var wg sync.WaitGroup

	for _, partition := range partitionedArr {
		wg.Add(1)
		go sortArray(partition, &wg)
	}

	wg.Wait()

	sortedArray := mergePartitions(partitionedArr)
	fmt.Println("Entire sorted array: ", sortedArray)
}

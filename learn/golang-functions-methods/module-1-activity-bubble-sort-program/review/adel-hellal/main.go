package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(sli []int) {
	for i := range sli {
		if i < len(sli)-1 && sli[i] > sli[i+1] {
			Swap(sli, i)
		}
	}
	if !isSorted(sli) {
		BubbleSort(sli)
	}
}
func isSorted(sli []int) bool {
	isSorted := true
	for i := range sli {
		if i < len(sli)-1 && sli[i] > sli[i+1] {
			isSorted = false
		}
	}
	return isSorted
}
func Swap(sli []int, index int) {
	temp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = temp

}

func main() {
	var sequenceInteger string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sequence of integer (eg : 4 2 1 30 5 6 0 50):")
	sequenceInteger, _ = reader.ReadString('\n')
	sequenceInteger = strings.Replace(sequenceInteger, "\n", "", 1)
	var sli []int
	strSli := strings.Split(sequenceInteger, " ")
	for i := range strSli {
		int_value, _ := strconv.ParseInt(strSli[i], 10, 32)
		sli = append(sli, int(int_value))
	}
	BubbleSort(sli)
	fmt.Println("sort result:", sli)
}

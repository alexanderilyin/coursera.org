package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// swap two consecutive integers in a slice. Assume the index passed is not out of bound
func Swap(sli []int, index int) {
	temp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = temp
}

func BubbleSort(sli []int) {
	n := len(sli)
	for i := n - 1; i > 0; i-- {
		for j := 0; j <= i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			} //if
		} // for j
	} // for i
}
func main() {

	var sli2 []int
	fmt.Printf("Enter up to 10 integers separated by exactly one space: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	text = strings.TrimSuffix(text, "\n")
	sli := strings.Split(text, " ")

	//	fmt.Printf("String before sort: %s len: %d \n", sli, len(sli))

	maxN := 10
	if len(sli) < 10 {
		maxN = len(sli)
	}
	for i := 0; i < maxN; i++ {
		item_i, _ := strconv.Atoi(sli[i])
		sli2 = append(sli2, item_i)
	}
	BubbleSort(sli2)

	//sli2 = strings.atoi(sli)
	fmt.Println(sli2)

}

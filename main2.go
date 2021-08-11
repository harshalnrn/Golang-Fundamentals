package main

import (
	"fmt"
)

func main() {
	intSlice := []int{9, 3, 6, 1, 2, 7}
	intSlice1 := []int{10, 3, 5, 1, 2, 7}
	intSlice2 := []int{0, 3, 9, 1, 2, 7}
	n := len(intSlice)
	n1 := len(intSlice1)
	n2 := len(intSlice2)

	//sort above with following sorting algorightsm in ascendin
	//bubble//insertion/selection/merge/quick

	//bubble
	//sorts array from right to left

	//outer shall deal with n elements to be sorted.
	for i := 1; i <= n-1; i++ {
		//inner shall deal with passes, where in each pass last element need not be sorted.
		for j := 0; j < n-i; j++ {
			if intSlice[j] > intSlice[j+1] {
				temp := intSlice[j]
				intSlice[j] = intSlice[j+1]
				intSlice[j+1] = temp
			}
		}
	}

	fmt.Println("bubble sort result")
	fmt.Println(intSlice)

	//selection
	//sort array from left to right
	//single swap per pass

	//selects min element index and then swaps finally only once per pass

	for i := 1; i <= n1-1; i++ {
		minIndex := i - 1
		for j := i; j <= n-1; j++ {
			//finding index of min element across this pass
			if intSlice1[j] < intSlice1[minIndex] {
				minIndex = j
			}
		}
		//swap element at minindex to ith position, thus sorting array from left to right

		temp := intSlice1[i-1]
		intSlice1[i-1] = intSlice1[minIndex]
		intSlice1[minIndex] = temp
	}

	fmt.Println("selection sort result")
	fmt.Println(intSlice1)

	//insertion sort

	//divide array into sorted and unsorted with left being sorted. and insert elements from unsorted into sorted itiratively

	//initially sorted array by deafult has just first element , and rest everything in unsorted
	for i := 1; i <= n2-1; i++ {

		for j := i; j >= 1; j-- {
			//check if last sorted element is > that first element of unsorted
			if intSlice2[j-1] > intSlice2[j] {
				//swap element from unsorted to sorted else break
				temp := intSlice2[j]
				intSlice2[j] = intSlice2[j-1]
				intSlice2[j-1] = temp
			} else {
				break
			}
		}
	}

	fmt.Println("insertion sort result")
	fmt.Println(intSlice2)

}

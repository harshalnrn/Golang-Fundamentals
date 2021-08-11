package main

import "fmt"

func main2() {

	sortedSlice := []int{1, 2, 3, 4, 5, 6}

	startindex := 0
	lastIndex := len(sortedSlice) - 1
	searchKey := 20

	for startindex <= lastIndex {

		midIndex := (startindex + lastIndex) / 2
		if sortedSlice[midIndex] == searchKey {
			//explcitly convert int to string before appending
			fmt.Println("binary search: key found at index" + fmt.Sprint(midIndex))
			break
		} else if sortedSlice[midIndex] > searchKey {

			lastIndex = midIndex - 1
		} else {

			startindex = midIndex + 1
		}
	}

	fmt.Println("search key " + fmt.Sprint(searchKey) + "not found in slice")

}

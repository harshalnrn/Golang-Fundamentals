package main

import "fmt"

func main() {

	//note slice is a reference type , and hence can have value=nil
	sortedSlice := []int{1, 2, 3, 4, 5, 6,20}
	startindex := 0
	lastIndex := len(sortedSlice) - 1
	searchKey := 20
	flag:=false

	for startindex <= lastIndex {

		midIndex := (startindex + lastIndex) / 2
		if sortedSlice[midIndex] == searchKey {
			//explcitly convert int to string before appending
			flag=true
			fmt.Println("binary search: key found at index: " + fmt.Sprint(midIndex))
			break
		} else if sortedSlice[midIndex] > searchKey {

			lastIndex = midIndex - 1
		} else {

			startindex = midIndex + 1
		}
	}

	if !flag {
	fmt.Println("search key " + fmt.Sprint(searchKey) + " not found in slice")
	}

}

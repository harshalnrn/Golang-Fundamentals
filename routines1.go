package main

import "fmt"

//channels are similar to blocking queue concept in java in context of 2 asychronous process

func main() {

	c1 := make(chan string) //assigned to 2 routines
	c2 := make(chan string) //assigned to 1 routine

	/* 	c1: 1,2,4
	   	c2: 3 */

	go func(c chan string) {

		c1 <- "child routine 1 completed" //enqueuing string into channel
		fmt.Println("routine 1 says bye")
	}(c1)

	go func(c chan string) {

		c1 <- "child rountine 2 completed"
		fmt.Println("routine 2 says bye")
	}(c1)

	go func(c chan string) {
		fmt.Println("before completion of routine 3")
		c2 <- "child routine 3 completed"
		fmt.Println("routine 3 says bye")

	}(c2)

	//channels behave like a queue here, where dequeue starts from earliest
	//example of controlling of order of routine/thread execution via channel comminication, as well as passing data between routines
	//  <- is dequequing channel
	m := <-c1
	fmt.Println(m)
	if m == "child routine 1 completed" {
		//create new routine and pass the data received in channel, to it
		go func(c chan string, message string) {
			fmt.Println("data received from previous routine: " + m)
			c1 <- "child routine 4 completed"
		}(c1, m)
	}

	fmt.Println(<-c2)
	fmt.Println(<-c1)
	fmt.Println(<-c1)

}

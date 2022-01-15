package main

import (
	"fmt"
	"net/http"
)

func main() {
	l := []string{"http://google.com", "http://golang.org", "http://amazon.com", "http://facebook.com"}
	//note: now when is channel required?
	c := make(chan string) //channel for string dataype //inbuilt make function
	//pass string message via channel to created routine and get back response
	for _, link := range l {
		//invoking function within separate go routine. go f(x)
		//This new goroutine will execute concurrently with the calling parent one. Thus Our two functions are running asynchronously in separate goroutines now
		//passing single channelt to multiple routines
		go getCall(link, c) // here channel should be pass by value / pass by reference ?
	}

	//create routine
	go func() {
		fmt.Println("create another routine with inline function declaration cum invocation")
	}()

	// lets understand channels , and how child routines can talk to its parent routine via channel
	//what are blocking channels
	//here channel behaves like a queue, where the earliest gets retreived/printed first
	fmt.Println(<-c) // channel blocks routine here, where routine pauses to consume received channel message
	//prnting each received message in the channel
	/* 	   	fmt.Println(<-c)
	   	   	fmt.Println(<-c)
	   	   	fmt.Println(<-c)  */

	//replace above with loop, where all messages get retreived from channel
	for i := 0; i < len(l); i++ {
		receivedMsg := <-c
		fmt.Println(receivedMsg) //print received response from channel
	}

	fmt.Println("main routine has come to end ")
}

func getCall(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		c <- "Might be down i think"
		return // note: empty return can be used in go, even for non returning methods
	}
	c <- "Yup its up" //write string into channel
}

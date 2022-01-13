/* package main // executable package

//non main package cannot be run(resusable package) . //only main package can be run (Executable package)
import "fmt"

//main function
//note: only "main" functona can be run..(the way java program searches for main method)
func main1() {
	fmt.Println("Hi there!")
}
*/

package main

import "fmt"

//datatypes, variables
/* 2 ways of LHS syntax
var <name> <type>
var:= <value>  //used for initializing local variables,  where var and type becoming optional, where golang figures it out auto. But not that golang is not dynamic like JS, hence type can be, as and when manipulated once set.  */

var data1 string
var data2 bool = true
var data3 int
var data4 float32
var data5 string = "john"
//data6:="amar" //Syntax error // Only local variables can be initialised with shorthand way using :=

//define function without return type
func main1() {
	//hellowordl function
	helloWorld()
	//learning methods and expressions
	fmt.Println(getMessage())

	result, message := getResult1(3, 4)
	fmt.Println(result, message)
	//control structures
	controlLadders()
	//creation and return size
	fmt.Println(createArray())
	fmt.Println(createSlice())

	//iterating blocks, on array and slice
	printArray()
	printSlice()

}

func helloWorld() {
	fmt.Println("hello world")
}

//function with return type
func getMessage() string {
	data6 := "ram"
	data1 = "helloWorld"
	addResult := getResult(4, 5) //number

	//note parsing between numbers and string in go.
	return data5 + data1 + data6 + fmt.Sprint(addResult)
}

//function taking parameter as well as has return types

func getResult(a int, b int) int {
	return a + b
}

//function returning multiple values
func getResult1(a int, b int) (int, string) {
	return a + b, "add result"
}

//do evaluation examples, where dealing with 'expressions' involving values of multiple data types.

func createArray() int {
	//intitializing array of strings
	var listOfBooks1 = [5]string{"fffs", "af", "bffh", "rwrw", "zczc"}
	listOfBooks2 := [5]string{"fffs", "aaf", "bffh", "rwrw", "zczc"} //making var optional

	fmt.Println(listOfBooks1[0])
	fmt.Println(listOfBooks2[1])

	return len(listOfBooks1)

}

func createSlice() int {
	var listOfBooks1 = []string{"fffs", "aaf", "bffh", "rwrw", "zczc"}
	listOfBooks2 := []string{"fffs", "aaf", "bffh", "rwrw", "zczc"} //making var optional

	fmt.Println(listOfBooks1[0])
	fmt.Println(listOfBooks2[1])

	listOfBooks3 := append(listOfBooks2, "vdvdv") //returns a new slice with appended element.// left to us whether to assign the new slice to new variable or existing

	fmt.Println(len(listOfBooks2))
	fmt.Println(len(listOfBooks3))
	return len(listOfBooks2)
}

// iterations,control strcutures on above data strcutures
func printArray() {
	//Go doesn't have variety of iterative lopp structures like while, do While, and has only 'for', making it syntax easy

	//note: go doesnt

	flag := true
	listOfBooks := [5]string{"fffs", "aaf", "bffh", "rwrw", "zczc"}

	//behaves like while loop
	for flag != false {
		fmt.Println("inside while like 'for' loop")
		flag = false
	}

	//basic for loop
	//note use of ;
	for i := 1; i <= 3; i++ {
		fmt.Println(listOfBooks[i])
	}
	//for each loop
	for _, currentElement := range listOfBooks {
		fmt.Println(currentElement)
	}

}

func printSlice() {
	listOfBooks := []string{"fffs", "aaf", "bffh", "rwrw", "zczc"}

	for i := 1; i <= 3; i++ {
		fmt.Println(listOfBooks[i])
	}

	for _, currentElement := range listOfBooks {
		fmt.Println(currentElement)
	}
}

//golang has similar ladders as other lang (if, if-else, if-elseif , switch)
func controlLadders() {
	a := 1
	//if block
	if a >= 1 {
		//formating string with dynamic values
		fmt.Println(fmt.Sprintf("value of a: %v", a))
	}

	//---if-else block-------------------
	if a >= 2 {
		//formating string with dynamic values
		fmt.Println(fmt.Sprintf("value of a: %v", a))

	} else {
		fmt.Println("invalid input")
	}

	//-----if-elseif-else block--------------------------
	if a >= 3 {
		//formating string with dynamic values
		fmt.Println(fmt.Sprintf("value of a: %v", a))

	} else if a > 4 {
		fmt.Println(a)
	} else if a > 5 {
		fmt.Println(a)
	} else {
		fmt.Println("invalid input")
	}

	//note above: in go, be cautious of code arrangement. (i.e if else)

	//switch : too many case conditions corresponding to single variable
	switch a {
	case 1:
		fmt.Print(a)
	case 2:
		fmt.Println(a)
	case 3:
		fmt.Println(a)
	default:
		fmt.Println("no cases met")
	}

}

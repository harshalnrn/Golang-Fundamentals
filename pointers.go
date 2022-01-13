package main
import "fmt"


type person struct{
	firstName string
	lastName string
	dob string
	phoneNumber string
}

func main() {

p:=person{firstName: "harshal", lastName:"narayan",dob:"23/06/1993",phoneNumber:"91-9887766"}
updateDetailsWithoutPointers(p) //pass by value
fmt.Println(fmt.Sprintf("updated struct after pass by value: %v",p))
updateDetailsWithPointers(&p)    //& represents pointer value, to be used for pointer type parameter.
fmt.Println(fmt.Sprintf("updated struct after pass by reference: %v",p))

}
//pass by value

func updateDetailsWithoutPointers(s person){

	s.firstName="amar"
}
//pass by reference

func updateDetailsWithPointers(s *person) {            //* represents pointer type.

	s.firstName="rajesh"

}



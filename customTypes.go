package main
import(
	"fmt"
) 
//only declaration-initialization statements can come outside functions

const num1,num2,num3=1,2,3    //initializing multiple constants in single line // datatypes not required here?
var data1,data2,data3 string        //declaring variables in single line of same datatype
//declaring custom types using interface and structs
type idRegisteration interface{
registerWithIdProof()
}



//implement above interface type with below child structs type
type AadharBasedRegisteration struct{

	name string
	aadharId string
	//dob time.Time
	//expiry time.time

}

type LicenseBasedRegisteration struct{
	name string
	licenseId string
	//dob time.Time
	//expiry time.Time

}




//-----------------------

func main(){
	// initialise struct variables and call implemented methods of interface of respective struct
	a:=AadharBasedRegisteration{aadharId: "1212121dadada"}
	l:=LicenseBasedRegisteration{licenseId:"12345"}
	a.registerWithIdProof()
	l.registerWithIdProof()
	//createAndPopulateMaps()
}

//Go is very similar to C due to presence of pointers, static typing and many other things. But Go is a modern language and so it has many new features baked in it.One notable feature is receiver functions.
//Receiver functions of structs for interface method implementation. With receiver functions you don’t have to mess around with classes or deal with inheritance

func (a AadharBasedRegisteration) registerWithIdProof(){

	fmt.Println("user registered with aadhar with aadhar id:", a.aadharId)
}


func (l LicenseBasedRegisteration) registerWithIdProof(){
	fmt.Println("user registered with licence with licence id:", l.licenseId)
}

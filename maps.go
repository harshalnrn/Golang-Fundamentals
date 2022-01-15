package main
import "fmt"


//map: list of (K,V) enteries
func main0(){
iterateMapEntries(createMaps())
}

func createMaps() map[int]string{
	m:=map[int]string{
		1:"harsh",
		2:"amar",
		3:"ajay",             //note , will follow even after the last element
	}
	fmt.Println(m)
	return m
}


func iterateMapEntries(m map[int]string){
	for id, name:=range m {
		fmt.Println(id,name)
	}
}

//empty interface{} is the "any" type, since all types implement the interface with no functions. Thus it should accept any type in golang (i.e similar to Object type in java)
//parent type to a child value
//understand type cast in line with use cases in java.
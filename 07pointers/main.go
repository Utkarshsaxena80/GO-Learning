package main 

import "fmt"

func main (){
	fmt.Println("hello");
	// var ptr *int
	// fmt.Println(ptr)
	myNumber:=23
	var ptr=&myNumber//referencing 
	fmt.Println(*ptr)
	fmt.Println(ptr)
	fmt.Println(*ptr*2)
}
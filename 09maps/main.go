package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["utkarsh"] = 99
	m["saxena"] = 100

	//orr

	val,exists := m["utkarsh"]
	if exists {
		fmt.Println(val)
	}else{
		fmt.Println("not found")
	}
	fmt.Println(len(m))
	//iterating over the map 
	for key,value:=range m{
		fmt.Println(key,value)
	}
	delete(m,"utkarsh")

	

}
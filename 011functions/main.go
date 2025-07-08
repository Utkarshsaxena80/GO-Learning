package main

import "fmt"

func main(){
	a:=10
	b:=10
	x:=add(a,b)
	fmt.Println(x)
}
func add(a int, b int)int{
	return a+b
}

//varidiac functions 

func sum(nums ...int) int{
	total:=0
	for _,num:=range nums{
		total+=num
	}
	return total
}
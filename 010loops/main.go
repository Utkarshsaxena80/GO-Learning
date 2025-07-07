package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}



	//iterate over a array or a map 
	numbers:=[]int{10,20,30,40}

	for index,value:=range numbers{//range gives both index and values 
		fmt.Println(index,value)
	}
	for _,value:=range numbers {//lets say we want only values 
		fmt.Println(value)
	}
	//no while loop in go 
	//only for is used 

	//impleamntaion of an exmaple of while loop 

	scanner:=bufio.NewScanner(os.Stdin)

	for{
		fmt.Println("enter text or type 'exit' to break")
		scanner.Scan()
		input := scanner.Text()

		if input=="exit"{
			break
		}
		fmt.Println("ypu entered : ",input )

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main (){
	fmt.Println("welcome to pizza app")
	reader:= bufio.NewReader(os.Stdin)
	//reads the data from terminal and stores in input 
	input,_:=reader.ReadString('\n')
	numRating,err:=strconv.ParseFloat(strings.TrimSpace(input),64);
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println("added one to your rating ...")
		fmt.Println(numRating+1)
	}
}
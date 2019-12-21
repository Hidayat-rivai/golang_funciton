package main

import "fmt"

func main(){
	fmt.Println(hello(10,5))
}

func hello(x int, y int)int{
	return x + y
}
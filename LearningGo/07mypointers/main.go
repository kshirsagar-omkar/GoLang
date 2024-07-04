package main

import (
	"fmt"
)

func main(){
	fmt.Println("Pointers");

	n := 10;



	var p *int = &n;

	fmt.Printf("Value of *p : %d\n%d",*p);
	fmt.Printf("Value of p : %p\n",p);



}
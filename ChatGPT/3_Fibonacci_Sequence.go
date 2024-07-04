/*
	Problem 3: Fibonacci Sequence
	Write a function that returns the first n numbers of the Fibonacci sequence as a slice.

	Function Signature:
	func fibonacci(n int) []int

	Example:
	fmt.Println(fibonacci(10))  // Output: [0 1 1 2 3 5 8 13 21 34]
*/

package main

import(
	"fmt"
)

func main(){
	fmt.Println("Problem 3: Fibonacci Sequence");
	fmt.Println(fibonacci(10));
}

func fibonacci(n int) []int{
	d1 := 0;
	d2 := 1;
	d3 := 0;
	var arr = []int{};
	//arr := make([]int,0)
	for i:=0; i<n; i++{

		arr = append(arr,d3);
		d1 = d2;
		d2 = d3;
		d3 = d1+d2;
	
	}
	return arr;
}
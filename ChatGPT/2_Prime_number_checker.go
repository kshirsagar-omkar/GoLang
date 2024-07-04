/*
	Problem 2: Prime Number Checker
	Write a function that checks if a given number is prime.

	Function Signature
	func isPrime(n int) bool
*/
package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
func main(){

	n,_ := takeInput();
	fmt.Println(isPrime(n));
	
}







func isPrime(n int64) bool{
	var i int64 = 2;
	for  i<n {
		if n%i==0{
			return false;
		}
		i++
	}
	return true;
}

func takeInput() (int64 ,error) {
	reader := bufio.NewReader(os.Stdin);
	
	fmt.Print("Enter Number : ");

	input,_ := reader.ReadString('\n');

	return (strconv.ParseInt(strings.TrimSpace(input),10,64));
}
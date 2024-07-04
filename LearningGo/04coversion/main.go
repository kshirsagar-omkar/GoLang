package main

import (

	"fmt"
	"os"
	"bufio"	
	"strconv"
	"strings"
)

func main(){

	fmt.Println("Welcome To Our Pizza App");
	fmt.Println("Please Rate Our Pizza Between 1-5 : ");

	reader := bufio.NewReader(os.Stdin);
	rating,_ := reader.ReadString('\n');

	fmt.Println("Thanks For Rating : ",rating);


	//numRating := rating +1;

	numRating,err := strconv.ParseFloat(strings.TrimSpace(rating),64);

	/*
	numRating,err := strconv.Atoi(rating);
	error : strconv.Atoi: parsing "1\n": invalid syntax
	*/

	if err != nil{
		fmt.Println(err);
	} else {
		fmt.Println("Added 1 to Radting : ",numRating+1);
	}


}
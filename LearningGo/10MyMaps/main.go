package main

import (
	"fmt"
)

func main(){

	fmt.Println("Maps");

	//count Frequency of each
	var arr = [20]int{1,2,4,23,1,2,4,2,3,244,1,1,2,2,3,2,4,2,5,1};
	fmt.Println("Array :",arr);

	freq := make(map[int]int);

	// for key,value := range freq{
	// 	freq[arr[key]] += 1;
	// }


	for i:=0; i<len(arr); i++{
		freq[arr[i]] += 1; 
	}

	delete(freq,244)
	

	for key,value := range freq{
		fmt.Println("No. :",key,"Freq :",value);
	}

}
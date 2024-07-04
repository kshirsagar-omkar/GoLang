/*
	Problem 7: Remove Duplicates from Slice
	Write a function that removes duplicate elements from a slice of integers.

	Function Signature:
	func removeDuplicates(slice []int) []int

	Example:
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 4, 4, 5}))  // Output: [1, 2, 3, 4, 5]
	fmt.Println(removeDuplicates([]int{10, 10, 20, 30}))       // Output: [10, 20, 30]

*/

package main

import(
	"fmt"
)

func main(){
	fmt.Println("7: Remove Duplicates from Slice");

	var arr = []int{1, 2, 2, 3, 4, 4, 5}

	//fmt.Println(removeDuplicates(arr));

	arr = removeDuplicates(arr);
	fmt.Println(arr);

	fmt.Println(removeDuplicates([]int{10, 10, 20, 30}))
}

func removeDuplicates(arr []int) []int{
	freq := make(map[int]int)

	for i := range arr{
		freq[arr[i]] += 1;
	}

	k := 0;
	for val,_:= range freq{
		arr[k] = val;
		k++;
	}

	arr = append(arr[:k])
	return arr;
}
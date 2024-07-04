package main
import (
	"fmt"
	"slices"
)

func main(){
	fmt.Println("My Slices");


	var slc = []int{11,22,33,44,55}

	fmt.Println("Slice slc :", slc);

	fmt.Println("slc[1:]",slc[1:]);

	fmt.Println("slc[:3]",slc[:3]);

	fmt.Println("slc[1:3]",slc[1:3]);


	fmt.Println("\n\nSize of slc :",len(slc),"\n");

	slc = append(slc,0);
	fmt.Println("Is Sort slc?",slices.IsSorted(slc));

	fmt.Println("\n\nSize of slc :",len(slc),"\n");

	slices.Sort(slc)
	fmt.Println("\nAfter Sorted :",slc)
	fmt.Println("Is Sort slc?",slices.IsSorted(slc));




	//Remove From Slices
	// 0,11,22,33,44,55
	i := 2;
	slc = append(slc[:i],slc[i+1:]...);
	fmt.Println("\n\nAfter Deleting value at index i :",i,"\nSlice is :",slc);
	fmt.Println("Size of slc :",len(slc),"\n"); 	







	arr := make([]int,2)
	fmt.Println("\n\nSize of arr :",len(arr),"\n");
	arr[0] = 11;
	arr[1] = 22;
	arr = append(arr,33);
	fmt.Println("\nSize of arr :",len(arr),"\n");

}
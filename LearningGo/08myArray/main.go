package main
import (
	"fmt"
)

func main(){
	fmt.Println("My Array");

	arr := [4]int{11,22,33,44}

	fmt.Println("Array length :",len(arr),"\nArray :",arr)

	i := 1
	//fmt.Println(arr[:i],arr[i+1:]);

	var slc []int = append(arr[:i],arr[i+1:]...)

	fmt.Println("Cutted Array of index i :",i,"array",slc)

}
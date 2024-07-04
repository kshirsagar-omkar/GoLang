/*
	Problem 6: Reverse a Slice
	Write a function that reverses the elements of a slice of integers.

	Function Signature:
	func reverseSlice(slice []int) []int

	Example:
	fmt.Println(reverseSlice([]int{1, 2, 3, 4, 5}))  // Output: [5, 4, 3, 2, 1]
	fmt.Println(reverseSlice([]int{10, 20, 30}))    // Output: [30, 20, 10]
*/

package main
import(
	"fmt"
)

func main(){
	fmt.Println("6: Reverse a Slice");
	var s = []int{1, 2, 3, 4, 5};
	fmt.Println(reverseSlice(s));

	fmt.Println(s);


}

func reverseSlice(arr []int) []int{
	var i int = 0;
	var j int = len(arr)-1;

	for i<j {	
		swapi(&arr[i],&arr[j])
		i++;
		j--;
	}
	return arr;
}

func swapi(a *int, b *int) {
	temp := *a;
	*a = *b;
	*b = temp;
}
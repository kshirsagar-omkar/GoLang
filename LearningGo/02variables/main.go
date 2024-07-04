package main
import "fmt"

var id = 101;
func main(){

	fmt.Println("Variables\n");

	var name string = "Omkar";
	fmt.Printf("Type %T : %s \n",name,name);

	var age uint8 = 18;
	fmt.Printf("Type %T : %d \n",age,age);

	fmt.Printf("Type %T : %d \n",id,id); 

	const PI = 3.14;
	fmt.Printf("Type %T : %f \n",PI,PI);



	var n1 = 10;
	var n2 = 20;
	fmt.Printf("\nAddition of %d+%d : %d",n1,n2,n1+n2); 
}
package main

import 
(
	"bufio"
	"fmt"
	"os"
)
func main() {

	welcome := "Welcome to User Input Program !";
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	fmt.Printf("Enter Message for output : ");

	//comma ok // comma err 

	//jeva error shi ky ghena dena ny  tr "_" vapraych, error konta? -> jr input detana ky problem ala tr to error // try catch nahi go madhe
	input, _ := reader.ReadString('\n');
	fmt.Println("Msg : ",input);


	//jr input ghetana ky problem ala tr "err" act as catch ani tyat thevta to error
	/*input, err := reader.ReadString('\n');
	fmt.Println("Msg : ",input);*/


	//jr input shi ghena dena nasla tr
	//_, err := reader.ReadString('\n');


}
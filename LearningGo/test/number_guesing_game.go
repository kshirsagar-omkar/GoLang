package main

import(

	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
	"math/rand"
)

func main() {

	rand.Seed(time.Now().UnixNano());
	var guessNum int64 = int64(rand.Intn(10));

	var chance uint8 = 1;

	fmt.Println("Welcome Ro Number Guessing Game !!");
	reader := bufio.NewReader(os.Stdin);

	
	fmt.Printf("Enter Number : ");

	StartAgain :
	input,err := reader.ReadString('\n');
	if err != nil {
		os.Exit(0)
	}

	num,err := strconv.ParseInt(strings.TrimSpace(input),10,64);
	if err != nil {
		os.Exit(0)
	}

	if num != guessNum {
		chance++;

		if num < guessNum {
			fmt.Println("Entered Number is smaller!!");
		}else {
			fmt.Println("Entered Number is Greater!!");
		}

		fmt.Printf("\nEnter Number Again : ");
		goto StartAgain; 
	}

	fmt.Printf("Guessed Right Number !!\nAfter %d chance\n",chance);

}
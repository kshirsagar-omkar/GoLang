package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("time package");
	presrntTime := time.Now();

	fmt.Println("Current Time :",presrntTime);

	fmt.Println("Formated Time :",presrntTime.Format("01-02-2006 15:04:05 Monday"));

}
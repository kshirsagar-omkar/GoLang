package main

import(
	"fmt"
	"sync"
)

func main(){
	fmt.Println("\t\t\t\t\t\t\t\t\t\t\t\tChannels\n");

							//Buffer channel
	myChn := make(chan int, 2);
	wg := &sync.WaitGroup{};

	// fmt.Println(<-myChn);
	// myChn <- 5;

	wg.Add(2);

	//Read Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch);
		//fmt.Println(<-ch);
		defer wg.Done();
	}(myChn,wg)

	//SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5;
		ch <- 6;
		close(ch);


		/*For Adding Ther will panic situation*/
		/*But For Readong There is No Panic Situation*/
		// close(ch);
		// /*read*/ch <- 5;
		// /*read*/ch <- 6;
		


		defer wg.Done();
	}(myChn,wg)


	wg.Wait();

}
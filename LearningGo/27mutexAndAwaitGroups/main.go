package main


import(
	"fmt"
	"sync"
)




// go run --race .

func main(){
	fmt.Println("Race Comdition = Youtube\n");

	wg := &sync.WaitGroup{};

	mut := &sync.RWMutex{};
	//There is another one named RWMutex read and write mutex
	//mut := &sync.RWMutex{};

	//mut.RLock()
	var score = []int{0};
	//mut.RUnlock()

	wg.Add(3);

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R");

		m.Lock();
		score = append(score,1);
		m.Unlock();

		defer wg.Done();
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R");

		m.Lock();
		score = append(score,2);
		m.Unlock();

		defer wg.Done();
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R");

		m.Lock();
		score = append(score,3);
		m.Unlock();

		defer wg.Done();
	}(wg, mut)


	wg.Add(1);
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R");

		m.RLock();
		fmt.Println(score)
		m.RUnlock();

		defer wg.Done();
	}(wg, mut)


	wg.Wait();





	//mut.RLock()
	fmt.Println(score);
	//mut.RUlnock()
}

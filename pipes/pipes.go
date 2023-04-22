package pipes

import (
	"fmt"
	"os"
	"strconv"
	utils "github.com/ptk-trindade/Distributed-Systems/utils"
)


func Pipe(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing parameters, add qty of numbers to generate")
		os.Exit(1)
	}

	qty, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: N value must be an integer")
		os.Exit(1)
	}
	
	pipeCh := make(chan int, 1)
	go Producer(pipeCh, qty)
	Consumer(pipeCh)

}

func Producer(pipeCh chan int, qty int) {
	var currentVal int
	for i:=0; i < qty; i++ {
		currentVal += utils.RandInt(100) + 1
		fmt.Println("p:", currentVal)
		pipeCh <- currentVal // insert into the pipe
	}

	pipeCh <- 0
}


func Consumer(pipeCh chan int) {
	currentVal := <- pipeCh // read from the pipe
	for currentVal != 0 {
		isPrime := utils.IsPrime(currentVal)

		if isPrime {
			fmt.Printf("c: %d is prime\n", currentVal)
		} else {
			fmt.Printf("c: %d is not prime\n", currentVal)
		}

		currentVal = <- pipeCh
	}
}
package main

import (
	// "syscall"
	"fmt"
	"os"
	"strconv"
)


func pipesPipe(args []string) {
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
	go producer(pipeCh, qty)
	consumer(pipeCh)

}

func producer(pipeCh chan int, qty int) {
	var currentVal int
	for i:=0; i < qty; i++ {
		currentVal += utilsRandInt(100) + 1
		fmt.Println("p:", currentVal)
		pipeCh <- currentVal // insert into the pipe
	}

	pipeCh <- 0
}


func consumer(pipeCh chan int) {
	currentVal := <- pipeCh // read from the pipe
	for currentVal != 0 {
		isPrime := utilsIsPrime(currentVal)

		if isPrime {
			fmt.Printf("c: %d is prime\n", currentVal)
		} else {
			fmt.Printf("c: %d is not prime\n", currentVal)
		}

		currentVal = <- pipeCh
	}
}
package pipes

import (
	"fmt"
	"os"
	"encoding/binary"
	"bytes"
	utils "github.com/ptk-trindade/Distributed-Systems/utils"
)


func Consumer(pipeReader *os.File) {
	
	currentVal := readFromPipe(pipeReader)
	for currentVal != int64(0) {
		isPrime := utils.IsPrime(int(currentVal))

		if isPrime {
			fmt.Printf("c: %d is prime\n", currentVal)
		} else {
			fmt.Printf("c: %d is not prime\n", currentVal)
		}

		currentVal = readFromPipe(pipeReader)
	}
}


func readFromPipe(pipeReader *os.File) int64 {
	byteSlice := make([]byte, 8)
	pipeReader.Read(byteSlice)
	
	var currentVal int64
	buf := bytes.NewReader(byteSlice)
	binary.Read(buf, binary.BigEndian, &currentVal)

	return currentVal
}
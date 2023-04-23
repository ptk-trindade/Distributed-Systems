package pipes

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
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

	pipeReader, pipeWriter, err := os.Pipe()

	if err != nil {
		fmt.Println("Error creating pipe:", err)
		os.Exit(1)
	}

	defer pipeReader.Close()
	defer pipeWriter.Close()

	pid, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if pid == 0 {
		Producer(pipeWriter, qty)
	} else if pid > 0 {
		Consumer(pipeReader)
	} else {
		fmt.Println("Error forking")
		os.Exit(1)
	}
}


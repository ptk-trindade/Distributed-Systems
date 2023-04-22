package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Error: Missing parameters, add parameter 'signals', 'pipes' or 'sockets'")
        os.Exit(1)
    }
    args := os.Args[1:]
    program, args := args[0], args[1:]

    switch program {
    case "signals":
        signalsSignal(args)
    case "pipes":
        pipesPipe(args)
    case "sockets":
        socketsSocket(args)

	default:
		fmt.Println("Invalid parameter, program must be 'signals', 'pipes' or 'sockets'")
	}

    fmt.Println("--- main end ---")
    return
}

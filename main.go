package main

import (
    "fmt"
    "os"
    signals "github.com/ptk-trindade/Distributed-Systems/signals"
    pipes "github.com/ptk-trindade/Distributed-Systems/pipes"
    sockets "github.com/ptk-trindade/Distributed-Systems/sockets"
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
        signals.Signal(args)
    case "pipes":
        pipes.Pipe(args)
    case "sockets":
        sockets.Socket(args)

	default:
		fmt.Println("Invalid parameter, program must be 'signals', 'pipes' or 'sockets'")
	}

    fmt.Println("--- main end ---")
    return
}

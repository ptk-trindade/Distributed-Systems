package signals

import (
    "fmt"
    "os"
)

// SIGNALS

/*
args = {send/receive [, ...]}
*/
func Signal(args []string) {
    if len(args) < 1 {
        fmt.Println("Error: Missing parameters, add parameter 'send' or 'receive'")
        os.Exit(1)
    }

    action, args := args[0], args[1:]
    switch action {
    case "send":
        Send(args)
    case "receive":
        Receive(args)

    default:
        fmt.Println("Error: Invalid parameter, action must be 'send' or 'receive'")
    }

    return
}



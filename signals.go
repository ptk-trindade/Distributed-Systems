package main // change to singal

import (
    "fmt"
    "strconv"
    "os"
    "os/signal"
    "syscall"
)

// SIGNALS

/*
args = {send/receive [, ...]}
*/
func signalsSignal(args []string) {
    if len(args) < 1 {
        fmt.Println("Error: Missing parameters, add parameter 'send' or 'receive'")
        os.Exit(1)
    }

    action, args := args[0], args[1:]
    switch action {
    case "send":
        signalsSend(args)
    case "receive":
        signalsReceive(args)

    default:
        fmt.Println("Error: Invalid parameter, action must be 'send' or 'receive'")
    }

    fmt.Println("--- Signal end ---")
    return
}


// RECEIVER
/*
args = {busy/blocking [, ...]}
*/
func signalsReceive(args []string) {
    if len(args) < 1 {
        fmt.Println("Error: Missing parameters, add parameter 'busy' or 'blocking'")
        os.Exit(1)
    }

    mode, args := args[0], args[1:]
    
    pid := syscall.Getpid()
    fmt.Printf("PID: %d, waiting for signal...\n", pid)
    
    sigCh := make(chan os.Signal, 1) // synchronous channel

    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPWR, nil)
    
    switch mode {
    case "busy":
        signalsReceiveBusyWait(sigCh)
    case "blocking":
        signalsReceiveBlocking(sigCh)
    default:
        fmt.Println("Error: Invalid parameter, mode must be 'busy' or 'blocking'")
    }
    
    
    return
}

func signalsReceiveBlocking(sigCh chan os.Signal) {
    fmt.Println("Blocking...")
    
    var signalReceived os.Signal
    for signalReceived != syscall.SIGINT { // signal 2 (0x2) - Interrupt from keyboard (Ctrl+C)
        signalReceived = <-sigCh
        fmt.Println("Received signal:", signalReceived)
    }

    fmt.Println("--- Blocking end ---")
    return
}


func signalsReceiveBusyWait(sigCh chan os.Signal) {
    fmt.Println("Busy waiting...")
    keepGoing := true
    for keepGoing {
        select {
            case signalReceived := <- sigCh:
                if signalReceived == syscall.SIGINT { // signal 2 (0x2) - Interrupt from keyboard (Ctrl+C)
                    fmt.Println("Received SIGINT")
                    keepGoing = false

                } else if signalReceived == syscall.SIGTERM { // signal 15 (0xf) - Termination signal
                    fmt.Println("Received SIGTERM")

                } else if signalReceived == syscall.SIGPWR { // signal 30 (0x1e) - System is shutting down
                    fmt.Println("Received SIGPWR")

                }

            default:
                fmt.Print(".") // no signal received
        }
        
    }

    fmt.Println("--- Busy wait end ---")
    return
}


// SENDER
/*
args = {pid, signal number}
*/
func signalsSend(args []string) {
    if len(args) < 2 {
        fmt.Println("Error: Missing parameters, add a PID and signal number to send")
        os.Exit(1)
    }

    pid, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Println("Error: Invalid PID, must be an integer")
        os.Exit(1)
    }

    sigNum, err := strconv.Atoi(args[1])
    if err != nil {
        fmt.Println("Error: Invalid signal number, must be an integer")
        os.Exit(1)
    }

    err = syscall.Kill(pid, syscall.Signal(sigNum)) // method ME

    /* method GPT
    cmd := exec.Command("kill", "-"+strconv.Itoa(sigNum), strconv.Itoa(pid))
    err = cmd.Run()
    */

    if err != nil {
        fmt.Println("Error sending signal:", err)
        os.Exit(1)
    }

    fmt.Printf("Sent signal %d to process with PID %d\n", sigNum, pid)
}

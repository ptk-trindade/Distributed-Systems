package signals

import (
    "fmt"
    "os"
	"os/signal"
    "syscall"
    "time"
)

/*
args = {busy/blocking [, ...]}
*/
func Receive(args []string) {
    if len(args) < 1 {
        fmt.Println("Error: Missing parameters, add parameter 'busy' or 'blocking'")
        os.Exit(1)
    }

    mode, args := args[0], args[1:]
    
    pid := syscall.Getpid()
    fmt.Printf("PID: %d, waiting for signal...\n", pid)
    
    sigCh := make(chan os.Signal, 1) // synchronous channel

    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPWR)
    
    switch mode {
    case "busy":
        ReceiveBusyWait(sigCh)
    case "blocking":
        ReceiveBlocking(sigCh)
    default:
        fmt.Println("Error: Invalid parameter, mode must be 'busy' or 'blocking'")
    }
    
    return
}


func ReceiveBlocking(sigCh chan os.Signal) {
    fmt.Println("Blocking...")
    
    var signalReceived os.Signal
    for signalReceived != syscall.SIGINT {
        signalReceived = <-sigCh

        if signalReceived == syscall.SIGINT { // signal 2 (0x2) - Interrupt from keyboard (Ctrl+C)
            fmt.Println("Received SIGINT")

        } else if signalReceived == syscall.SIGTERM { // signal 15 (0xf) - Termination signal
            fmt.Println("Received SIGTERM")

        } else if signalReceived == syscall.SIGPWR { // signal 30 (0x1e) - System is shutting down
            fmt.Println("Received SIGPWR")
            
        } else {
            fmt.Println("Received unknown signal:", signalReceived)
            
        }
    }

    return
}


func ReceiveBusyWait(sigCh chan os.Signal) {
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
                    
                } else {
                    fmt.Println("Received unknown signal:", signalReceived)

                }

            default:
                time.Sleep(1 * time.Second)
                fmt.Print(".") // no signal received
        }
        
    }

    return
}

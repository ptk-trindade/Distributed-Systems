package signals

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "syscall"
)


func Send(args []string) {
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

    err = syscall.Kill(pid, syscall.Signal(sigNum))

    if err != nil {
        fmt.Println("Error sending signal:", err)
        os.Exit(1)
    }

    log.Printf("Sent signal %d to process with PID %d\n", sigNum, pid)
}

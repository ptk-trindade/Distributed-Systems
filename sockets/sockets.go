package sockets

import (
    "fmt"
    "net"
	"os"
	"strconv"
	"encoding/binary"
	"bytes"
)


func Socket(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing parameters, add parameter 'client' or 'server'")
		os.Exit(1)
	}

	action, args := args[0], args[1:]
	switch action {
	case "client":
		Client(args)
	case "server":
		Server()
	}

}
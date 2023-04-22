package sockets

import (
    "fmt"
    "net"
	"os"
	"strconv"
	"encoding/binary"
	"bytes"
)

func Server() {
    // listen on port 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server started, listening on port 8080")

	closeServer := make(chan bool, 1)
    for {
        // wait for a connection
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
			fmt.Println(err)
            continue
        }

        go handleConnection(conn, closeServer)
    }
}


// ----- LOCAL FUNCTIONS ------

func handleConnection(conn net.Conn, closeServer chan bool) {
    defer conn.Close()
	defer fmt.Println("Closing connection...")
	
	for {
		valReceived, err := readFromClient(conn)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			continue
		}
		
		if valReceived <= int64(0) { // close connection
			if valReceived == int64(-1) { // close server
				fmt.Println("Closing server...")
				os.Exit(0)
			}
			return
		}
		
		isPrime := utilsIsPrime(int(valReceived))
		var response string
		if isPrime {
			response = "Is prime    "
		} else {
			response = "Is not prime"
		}

		err = sendToClient(conn, response)
		if err != nil {
			fmt.Println("Error sending to client:", err)
			continue
		}
	}

}


func readFromClient(conn net.Conn) (int64, error) {
	byteSlice := make([]byte, 8)
	_, err := conn.Read(byteSlice) // first return value is qty of bytesRead (will be 8)
	if err != nil {
		return 0, err
	}

	var valReceived int64
	buf := bytes.NewReader(byteSlice)
	binary.Read(buf, binary.BigEndian, &valReceived) // bufReader[:bytesRead] (bytesRead is 8)
	fmt.Printf("Received value: %d\n", valReceived) // TEMP

	return valReceived, nil
}


func sendToClient(conn net.Conn, response string) error {
	if len(response) != 12 {
		return fmt.Errorf("Error: text must be 12 characters long")
	}

	var buf bytes.Buffer
	buf.WriteString(response)
	_, err := conn.Write(buf.Bytes())

	return err
}
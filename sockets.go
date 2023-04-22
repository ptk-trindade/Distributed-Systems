package main

import (
    "fmt"
    "net"
	"os"
	"strconv"
	"encoding/binary"
	"bytes"
)


func socketsSocket(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing parameters, add parameter 'client' or 'server'")
		os.Exit(1)
	}

	action, args := args[0], args[1:]
	switch action {
	case "client":
		socketsClient(args)
	case "server":
		socketsServer()
	}

}

// CLIENT
func socketsClient(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing parameters, add parameter 'N'")
		os.Exit(1)
	}

	qty, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: N value must be an integer")
		os.Exit(1)
	}

	var closingValue int64 = 0
	if len(args) > 1 && args[1] == "close"{
		closingValue = -1
	}


    // connect to the server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
		fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()
	
	
	current_n := int64(utilsRandInt(100) + 1)
	for i:=0; i < qty; i++ {
		fmt.Println("Value sent:", current_n)
		err := sendToServer(conn, current_n)
		if err != nil {
			fmt.Println("Error sending message to server:", err)
			return
		}
		

		current_n += int64(utilsRandInt(100) + 1)

        // read the response from the server
        responseStr, err := readFromServer(conn)
        if err != nil {
            fmt.Println("Error reading response from server:", err)
            return
        }

        fmt.Println("Server response:", responseStr)
    }

	sendToServer(conn, closingValue)
}


func sendToServer(conn net.Conn, value int64) error {
	// convert the integer to a byte slice and send it to the server
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, value)

	_, err := conn.Write(buf.Bytes())

	return err
}


func readFromServer(conn net.Conn) (string, error) {
	// read the response from the server
	byteSlice := make([]byte, 12)
	_, err := conn.Read(byteSlice)
	if err != nil {
		return "", err
	}

	return string(byteSlice), err
}



// SERVER

func socketsServer() {
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


func handleConnection(conn net.Conn, closeServer chan bool) {
    defer conn.Close()
	defer fmt.Println("Closing connection...")
	
	for {
		byteSlice := make([]byte, 8)
		_, err := conn.Read(byteSlice) // first return value is bytesRead (will be 8)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		var valReceived int64
		bufReader := bytes.NewReader(byteSlice)
		binary.Read(bufReader, binary.BigEndian, &valReceived) // buf[:bytesRead] (bytesRead is 8)
		fmt.Printf("Received value: %d\n", valReceived) // TEMP
		
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

		var buf bytes.Buffer
		buf.WriteString(response)
		conn.Write(buf.Bytes())
	}

}
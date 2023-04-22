package sockets

import (
    "fmt"
    "net"
	"os"
	"strconv"
	"encoding/binary"
	"bytes"
)

func Client(args []string) {
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



// ----- LOCAL FUNCTIONS ------

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




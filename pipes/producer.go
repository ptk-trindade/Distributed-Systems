package pipes

import (
	"fmt"
	"os"
	"encoding/binary"
	"bytes"
	utils "github.com/ptk-trindade/Distributed-Systems/utils"
)

func Producer(pipeWriter *os.File, qty int) {
	var currentVal int
	for i:=0; i < qty; i++ {
		currentVal += utils.RandInt(100) + 1
		fmt.Println("p:", currentVal)

		writeOnPipe(pipeWriter, currentVal)
	}

	writeOnPipe(pipeWriter, 0)
}

func writeOnPipe(pipeWriter *os.File, currentVal int) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, int64(currentVal))
	
	pipeWriter.Write(buf.Bytes())
}
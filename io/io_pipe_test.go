package main

import (
	"bufio"
	"fmt"
	"io"
	"testing"
	"time"
)

func TestIoPipe(t *testing.T) {
	reader, writer := io.Pipe()

	go IoPipeRead(reader)

	writer.Write([]byte("hello"))
	writer.Write([]byte("world"))

	writer.Close()

	time.Sleep(1 * time.Second)
}

func IoPipeRead(reader *io.PipeReader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	reader.Close()
}
package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        "COM1",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	// Write 4 bytes to the port.
	b := []byte{0x02, 0x48, 0x45, 0x4C, 0x4C, 0x4F, 0x0A, 0x03}
	n, err := port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
}

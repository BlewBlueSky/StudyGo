package main

import (
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        "COM7",
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

	// Write 9 bytes to the port.
	arr := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	n, err := port.Write(arr)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	log.Printf("Wrote %v bytes:%q\n", n, arr[:n])
	fmt.Printf("Wrote %v bytes:%q\n", n, arr[:n])

	// Read 9 bytes from the port
	buf := make([]byte, 128)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read %v bytes:%q\n", n, buf[:n])
	fmt.Printf("Read %v bytes:%q\n", n, buf[:n])

}

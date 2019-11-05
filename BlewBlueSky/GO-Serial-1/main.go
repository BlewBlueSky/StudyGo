package main


import "fmt"
import "log"
import "github.com/jacobsa/go-serial/serial"


func main(){
	// Set up options.
	options := serial.OpenOptions{
	PortName: "COM7",
	BaudRate: 9600,
	DataBits: 8,
	StopBits: 1,
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
	b := []byte{0x01, 0x00, 0x00, 0x00,0x00, 0x00, 0x00, 0x00,0x01}
	n, err := port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}	
	fmt.Println("Wrote", n, "bytes.")

	//
	buf := make([]byte, 128)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q:%v\n", buf[:n],n)
}

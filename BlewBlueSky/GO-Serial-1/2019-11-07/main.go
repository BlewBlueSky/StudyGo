package main

import (

	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

var (
	trinamicShop  = []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	arr           []byte
	buf           []byte
	bufn          int
	i             int32
	returnAddress byte
	returnStatus  byte
	returnValue   int
	address       byte
	ok1			  int8
)

const (
	//define axis
	axis0 = 0
	axis1 = 1
	axis2 = 2
	axis3 = 3
	axis4 = 4
	axis5 = 5

	//Opcodes of all TMCL commands that can be used in direct mode

	tmclROR  = 1
	tmclROL  = 2
	tmclMST  = 3
	tmclMVP  = 4
	tmclSAP  = 5
	tmclGAP  = 6
	tmclSTAP = 7
	tmclRSAP = 8
	tmclSGP  = 9
	tmclGGP  = 10
	tmclSTGP = 11
	tmclRSGP = 12
	tmclRFS  = 13
	tmclSIO  = 14
	tmclGIO  = 15
	tmclSCO  = 30
	tmclGCO  = 31
	tmclCCO  = 32

	//Opcodes of TMCL control functions (to be used to run or abort a TMCL program in the module)

	tmclApplStop  = 128
	tmclApplRun   = 129
	tmclApplReset = 131

	//Options for MVP commandds

	mvpABS   = 0
	mvpREL   = 1
	mvpCOORD = 2

	//Options for RFS command

	rfsSTART  = 0
	rfsSTOP   = 1
	rfsSTATUS = 2

	//Result codes for GetResult

	tmclReesultOK           = 0
	tmclReesultNotReady     = 1
	tmclReesulChecksumERROR = 2
)

//COM INIT

func comInit(arr []byte, com string) ([]byte, int) {

	options := serial.OpenOptions{
		PortName:        com,
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

	n, err := port.Write(arr)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	log.Println("Wrote", n, "bytes:", arr[:n])

	// Read 9 bytes from the port
	buf := make([]byte, 128)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read", n, "bytes:", buf[:n])
	return buf, n
}

//Send CMD

func sendCmd(COM string, Address uint8, Command uint8, Type uint8, Motor uint8, Value int32) ([]byte, int) {

	var txBuffer = []byte{Address, Command, Type, Motor, byte(Value >> 24), byte(Value >> 16), byte(Value >> 8), byte(Value & 0xff), 0}
	var bufn int
	/*txBuffer[0]=Address
	txBuffer[1]=Command
	txBuffer[2]=Type
	txBuffer[3]=Motor
	txBuffer[4]=byte(Value >> 24)
	txBuffer[5]=byte(Value >> 16)
	txBuffer[6]=byte(Value >> 8)
	txBuffer[7]=byte(Value & 0xff)
	txBuffer[8]=0*/
	for i = 0; i < 8; i++ {
		txBuffer[8] += txBuffer[i]
	}

	//Send the datagram
	txBuffer, bufn = comInit(txBuffer, "COM7")
	//fmt.Println("final:",txBuffer[:bufn])
	return txBuffer, bufn
}

func getResult(rxBuffer []byte) int8 {
	var checksum byte = 0
	if len(rxBuffer) >= 8 {

		for i = 0; i < 8; i++ {
			checksum += rxBuffer[i]
		}
			if checksum != rxBuffer[8] {
				fmt.Println("TMCL Reesul Checksum ERROR")
				return tmclReesulChecksumERROR
			}
			returnAddress = rxBuffer[0]
			returnStatus = rxBuffer[2]

			returnValue = (int(rxBuffer[4])<<24)|(int(rxBuffer[5])<<16)|(int(rxBuffer[6])<<8)|int(rxBuffer[7])
		
	}else {
		fmt.Println("TMCL_Reesult_Not_Ready")
		return tmclReesultNotReady
		
	}
	fmt.Println("TMCL_Reesult_OK")
	return tmclReesultOK
}

func main() {

	buf, bufn = sendCmd("COM7", 1, tmclROR, 0, 0, 100)
	getResult(buf)
	//buf,bufn =comInit(buf,"COM7")
	fmt.Println(ok1,returnAddress,returnValue,returnStatus)

}

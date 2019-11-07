package main

import (
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)


	
var (
	trinamicShop = []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	arr []byte
	buf [9]byte
	bufn int
)

const(
	//define axis
	axis0 = 0
	axis1 = 1
	axis2 = 2
	axis3 = 3
	axis4 = 4
	axis5 = 5
	

	//Opcodes of all TMCL commands that can be used in direct mode

	tmclROR = 1
	tmclROL = 2
	tmclMST = 3
	tmclMVP = 4
	tmclSAP = 5
	tmclGAP = 6
	tmclSTAP = 7
	tmclRSAP = 8
	tmclSGP = 9
	tmclGGP = 10
	tmclSTGP = 11
	tmclRSGP = 12
	tmclRFS = 13
	tmclSIO = 14
	tmclGIO = 15
	tmclSCO = 30
	tmclGCO = 31
	tmclCCO = 32

	//Opcodes of TMCL control functions (to be used to run or abort a TMCL program in the module)

	tmclApplStop = 128
	tmclApplRun = 129
	tmclApplReset = 131

	//Options for MVP commandds

	mvpABS = 0
	mvpREL = 1
	mvpCOORD = 2

	//Options for RFS command

	rfsSTART = 0
	rfsSTOP = 1
	rfsSTATUS = 2



	//Result codes for GetResult

	tmclReesultOK = 0
	tmclReesultNotReady = 1
	tmclReesulChecksumERROR = 2
)

func comInit(arr []byte,com string )([]byte,int)  {

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

	log.Printf("Wrote %v bytes:%q\n", n, arr[:n])

	// Read 9 bytes from the port
	buf := make([]byte, 128)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read %v bytes:%q\n", n, buf[:n])
	return buf, n
}

func sendCmd(COM string, Address uint8, Command uint8, Type uint8, Motor uint8, Value int32)([9]byte){
	
	var TxBuffer [9]byte
	var i int32

	TxBuffer[0]=Address
	TxBuffer[1]=Command
	TxBuffer[2]=Type
	TxBuffer[3]=Motor
	TxBuffer[4]=byte(Value >> 24)
	TxBuffer[5]=byte(Value >> 16)
	TxBuffer[6]=byte(Value >> 8)
	TxBuffer[7]=byte(Value & 0xff)
	TxBuffer[8]=0
	for i=0; i<8; i++{
		TxBuffer[8]+=TxBuffer[i]
	}
		

	//Send the datagram
	//TxBuffer,bufn = comInit(TxBuffer,"COM7")
	fmt.Printf("final:%q\n",  TxBuffer[:9])
	return TxBuffer
}


func main() {

	buf = sendCmd("COM7",1,tmclROR,0,0,100)
	fmt.Println(buf)
	//buf,bufn =comInit(buf,"COM7")
	
					
}

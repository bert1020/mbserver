package mbserver

import (
	"fmt"
	"testing"
)

func TestCRC(t *testing.T) {
	got := crcModbus([]byte{0x55, 0x12, 0x34, 0x03, 0x04, 0x1E})
	fmt.Printf("<UNK>: % X \n", got)

}

package mbserver

import (
	"testing"
)

func TestCRC(t *testing.T) {
	//got := crcModbus([]byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x01})
	//fmt.Printf("<UNK>: %04X \n", got)
	//expect := 0x840A
	//if !isEqual(expect, got) {
	//	t.Errorf("expected %x, got %x", expect, got)
	//}
	got := crcModbus([]byte{0x01, 0x04, 0x02, 0xFF, 0xFF})
	expect := 0x80B8
	if !isEqual(expect, got) {
		t.Errorf("expected %x, got %x", expect, got)
	}
}

package mbserver

import (
	"log"
	"testing"
	"time"
)

func TestNewRTUFrame(t *testing.T) {
	for {
		_, err := NewRTUFrame([]byte{0xFF, 0x03, 0x01, 00, 00, 00, 0x0A, 0xBD, 0xEF})
		if !isEqual(nil, err) {
			log.Printf("NewRTUFrame failed: %v", err)
			//t.Fatalf("expected %v, got %v", nil, err)
		}
		time.Sleep(time.Millisecond * 1000)
	}

}

func TestNewRTUFrameShortPacket(t *testing.T) {
	_, err := NewRTUFrame([]byte{0x01, 0x04, 0xFF, 0xFF})
	if err == nil {
		t.Fatalf("expected error not nil, got %v", err)
	}
}

func TestNewRTUFrameBadCRC(t *testing.T) {
	// Bad CRC: 0x81 (should be 0x80)
	_, err := NewRTUFrame([]byte{0x01, 0x04, 0x02, 0xFF, 0xFF, 0xB8, 0x81})
	if err == nil {
		t.Fatalf("expected error not nil, got %v", err)
	}
}

func TestRTUFrameBytes(t *testing.T) {
	frame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(4),
		Data:     []byte{0x02, 0xff, 0xff},
	}

	got := frame.Bytes()
	expect := []byte{0x01, 0x04, 0x02, 0xFF, 0xFF, 0xB8, 0x80}
	if !isEqual(expect, got) {
		t.Errorf("expected %v, got %v", expect, got)
	}
}

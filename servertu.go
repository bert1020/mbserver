package mbserver

import (
	"io"
	"log"

	"github.com/goburrow/serial"
)

// ListenRTU starts the Modbus server listening to a serial device.
// For example:  err := s.ListenRTU(&serial.Config{Address: "/dev/ttyUSB0"})
func (s *Server) ListenRTU(serialConfig *serial.Config) (err error) {
	port, err := serial.Open(serialConfig)
	if err != nil {
		log.Fatalf("failed to open %s: %v\n", serialConfig.Address, err)
	}
	s.ports = append(s.ports, port)

	s.portsWG.Add(1)
	go func() {
		defer s.portsWG.Done()
		s.acceptSerialRequests(port)
	}()

	return err
}

var (
	accumulatedData []byte
	buffer          = make([]byte, 256) // 缓冲区大小需≥最大报文长度
)

func (s *Server) acceptSerialRequests(port serial.Port) {
SkipFrameError:
	for {
		select {
		case <-s.portsCloseChan:
			return
		default:
		}

		bytesRead, err := port.Read(buffer)
		//fmt.Printf("Read %d bytes, bytes :% X\n", bytesRead, buffer[:bytesRead])
		if bytesRead == 0 {
			continue
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("serial read error %v\n", err)
				continue
			}
		}

		accumulatedData = append(accumulatedData, buffer[:bytesRead]...)
		if bytesRead >= 5 {

			frame, err := NewRTUFrame(accumulatedData)
			if err != nil {
				continue SkipFrameError
				//return
			}
			accumulatedData = nil

			request := &Request{port, frame}

			s.requestChan <- request
		} else {

		}
	}
}

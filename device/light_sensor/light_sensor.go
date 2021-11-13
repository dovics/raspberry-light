package light_sensor

import (
	"bufio"
	"github.com/tarm/serial"
)

type LightSensor struct {
	port *serial.Port
	buf *bufio.Reader
}

func Connect(port *serial.Port) *LightSensor {
	return &LightSensor{
		port: port,
		buf: bufio.NewReader(port),
	}
}

func (s *LightSensor) ReadLine() (string, error) {
	line, _, err := s.buf.ReadLine()
	return string(line), err
}

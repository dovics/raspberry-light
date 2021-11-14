package light_sensor

import (
	"testing"
	"time"

	"github.com/tarm/serial"
)

func TestLightSensor(t *testing.T) {
	c := &serial.Config{Name: "/dev/serial0", Baud: 9600, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	sensor := Connect(s)
	for {
		data, err := sensor.ReadLine()
		if err != nil {
			t.Error(err)
		}

		t.Log(data)
	}

}

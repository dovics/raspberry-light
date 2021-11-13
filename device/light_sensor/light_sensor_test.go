package light_sensor

import (
	"time"
	"testing"

	"github.com/tarm/serial"
)

func TestLightSensor(t *testing.T) {
	c := &serial.Config{Name: "/dev/serial0", Baud: 9600, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	sensor := Connect(c)
	for {
		data, err := sensor.ReadLine()
		if err != nil {
			LogError(err)
		}

		t.Log(data)
	}

}

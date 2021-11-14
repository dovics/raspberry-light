package main

import (
	"time"

	"github.com/tarm/serial"

	"github.com/dovics/raspberry-light/device/light_sensor"
	"github.com/dovics/raspberry-light/operator"
)

func main() {
	c := &serial.Config{Name: "/dev/serial0", Baud: 9600, ReadTimeout: time.Second * 5}

	sensor, err := light_sensor.ConnectBySerial(c)
	if err != nil {
		panic(err)
	}
	operator.NewLightOperator(sensor)

	for{}
}

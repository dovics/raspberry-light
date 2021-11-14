package main

import (
	"time"

	"github.com/tarm/serial"

	"github.com/dovics/raspberry-light/device/light_sensor"
	"github.com/dovics/raspberry-light/device/relay"
	"github.com/dovics/raspberry-light/exporter"
	"github.com/dovics/raspberry-light/operator"
	"github.com/dovics/raspberry-light/reporter"
)

func main() {
	c := &serial.Config{Name: "/dev/serial0", Baud: 9600, ReadTimeout: time.Second * 5}

	sensor, err := light_sensor.ConnectBySerial(c)
	if err != nil {
		panic(err)
	}
	lightOperator := operator.NewLightOperator(sensor)

	r := relay.Connect()
	relayOperator := operator.NewRelayOperator(r)
	reporter := reporter.New("")

	reporter.SetTrigger(lightOperator)
	reporter.Register("relay", func() (interface{}, error) {
		if lightOperator.IsOpen() {
			relayOperator.Open()
		} else {
			relayOperator.Close()
		}

		return "Success", nil
	})

	go reporter.Run()

	exporter := exporter.NewExporter()
	exporter.Register("light", lightOperator.QueryLight)
	exporter.Register("switch", relayOperator.Switch)

	exporter.Run()
}

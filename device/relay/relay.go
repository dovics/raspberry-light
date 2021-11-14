package relay

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type Relay struct {
	ina rpio.Pin
	inb rpio.Pin
}

func Connect() *Relay {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	ina := rpio.Pin(17)
	inb := rpio.Pin(27)

	ina.Output()
	inb.Output()

	return &Relay{
		ina,
		inb,
	}
}

func (r *Relay) OpenA() {
	r.ina.High()
}

func (r *Relay) OpenB() {
	r.inb.High()
}

func (r *Relay) IsOpenA() bool {
	return r.ina.Read() == rpio.High
}

func (r *Relay) IsOpenB() bool {
	return r.inb.Read() == rpio.High
}

func (r *Relay) CloseA() {
	r.ina.Low()
}

func (r *Relay) CloseB() {
	r.inb.Low()
}

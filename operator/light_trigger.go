package operator

import (
	"github.com/dovics/raspberry-light/device/light_sensor"
	"github.com/dovics/raspberry-light/util/log"
)

type LightTrigger struct {
	sensor light_sensor.LightSensor
	ch     chan struct{}
}

func NewLightTrigger(sensor light_sensor.LightSensor) *LightTrigger {
	trigger := &LightTrigger{
		sensor: sensor,
		ch:     make(chan struct{}),
	}

	go trigger.run()
	return trigger
}

func (t *LightTrigger) Chan() <-chan struct{} {
	return t.ch
}

func (t *LightTrigger) run() {
	for {
		line, err := t.sensor.ReadLine()
		if err != nil {
			log.Info(err)
		}
		if ParseData(line) < 1000 {
			select {
			case t.ch <- struct{}{}:
			default:
			}
		}
	}
}

func ParseData(line string) int {
	return 0
}

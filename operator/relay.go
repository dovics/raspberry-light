package operator

import "github.com/dovics/raspberry-light/device/relay"

type RelayOperator struct {
	relay *relay.Relay
}

func NewRelayOperator(r *relay.Relay) *RelayOperator {
	return &RelayOperator{
		relay: r,
	}
}

func (o *RelayOperator) Switch() (interface{}, error) {
	if o.relay.IsOpenA() {
		o.relay.CloseA()
	} else {
		o.relay.OpenA()
	}

	return "Success", nil
}

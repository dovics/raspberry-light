package relay

import (
	"testing"
	"time"
)

func TestRelay(t *testing.T) {
	relay := Connect()
	relay.OpenA()
	relay.OpenB()

	t.Log("Open Success")
	time.Sleep(time.Second * 3)

	relay.CloseA()
	relay.CloseB()

	t.Log("Close Success")
}

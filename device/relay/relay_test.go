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

func TestRelayIsOpen(t *testing.T) {
	relay := Connect()
	relay.OpenA()

	if !relay.IsOpenA() {
		t.Error("A should be open")
	}

	if relay.IsOpenB() {
		t.Error("B should be close")
	}
}

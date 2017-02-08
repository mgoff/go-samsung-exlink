package exlink

import (
	"testing"
	"time"
)

// udpate devicePath with your serial port path before running 'go test'
const (
	devicePath = "/dev/tty.KeySerial1"
)

// sleep in between test commands so you can watch the results on your TV
func TestSamsungExlink(t *testing.T) {

	// open the serial connection to the EX-Link device
	device, err := Open(devicePath)
	if err != nil {
		t.Error(err)
	}

	// turn on the EX-Link device
	if device.PowerOn() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// volume up
	if device.VolumeUp() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// volume down
	if device.VolumeDown() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// speaker off
	if device.SpeakerOff() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// speaker on
	if device.SpeakerOn() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// switch to HDMI 1
	if device.SourceHDMI1() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// switch to HDMI 2
	if device.SourceHDMI2() != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)

	// turn off the EX-Link device
	if device.PowerOff() != nil {
		t.Error(err)
	}

	// close the serial connection to the EX-Link device
	if device.Close() != nil {
		t.Error(err)
	}
}

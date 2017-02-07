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

	// open the connection to the EX-Link device
	device, err := Open(devicePath)
	if err != nil {
		t.Error(err)
		return
	}

	// turn on the EX-Link device
	err = device.PowerOn()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// volume up
	err = device.VolumeUp()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// volume down
	err = device.VolumeDown()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// speaker off
	err = device.SpeakerOff()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// speaker on
	err = device.SpeakerOn()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// switch to HDMI 1
	err = device.SourceHDMI1()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// switch to HDMI 2
	err = device.SourceHDMI2()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// turn off the EX-Link device
	err = device.PowerOff()
	if err != nil {
		t.Error(err)
		return
	}

	// close the connection to the EX-Link device
	err = device.Close()
	if err != nil {
		t.Error(err)
	}
}

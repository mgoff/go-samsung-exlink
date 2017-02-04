package exlink

import (
	"encoding/hex"
	"github.com/jacobsa/go-serial/serial"
	"io"
)

// hex codes from the EX-Link protocol
const (
	powerOn		= "082200000002d6"
	powerOff	= "082200000001d5"
	volumeUp	= "082201000100d4"
	volumeDown	= "082201000200d3"
	speakerOn	= "08220c060000c4"
	speakerOff	= "08220c060001c3"
	hdmi1		= "08220a000500c7"
	hdmi2		= "08220a000501c6"
)

// serial port parameters for EX-Link devices
const (
	baudRate	= 9600
	dataBits	= 8
	stopBits	= 1
	minReadSize	= 4
)

// struct used to hold our serial connection
type Exlink struct {
	port io.ReadWriteCloser
}

// open a serial connection to the EX-Link device
func Open(device string) (Exlink, error) {

	// setup serial port options
	options := serial.OpenOptions{
		PortName: device,
		BaudRate: baudRate,
		DataBits: dataBits,
		StopBits: stopBits,
		MinimumReadSize: minReadSize,
  	}

	// open the port
	port, err := serial.Open(options)
	if err != nil {
		return Exlink{}, err
	} else {
		exlink := Exlink{port:port}
		return exlink, nil
	}
}

// send a hex code to the EX-Link device
func (e Exlink) Send(code string) error {

	// convert the hex string into bytes
	b, err := hex.DecodeString(code)
	if err != nil {
		return err
	}

	// write the bytes
	_, err = e.port.Write(b)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// close the serial connection to the EX-Link device
func (e Exlink) Close() error {
	err := e.port.Close()
	if err != nil {
		return err
	} else {
		e.port = nil
		return nil
	}
}

// power on the EX-Link device
func (e Exlink) PowerOn() error {
	err := e.Send(powerOn)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// power off the EX-Link device
func (e Exlink) PowerOff() error {
	err := e.Send(powerOff)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// switch to source input HDMI1
func (e Exlink) SourceHDMI1() error {
	err := e.Send(hdmi1)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// switch to source input HDMI2
func (e Exlink) SourceHDMI2() error {
	err := e.Send(hdmi2)
	if err != nil {
		return err
	} else {
		return nil
	}
}
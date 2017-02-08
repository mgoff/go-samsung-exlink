# go-samsung-exlink
A Go library for controlling Samsung TVs through the EX-Link serial port included on most models.

Samsung EX-Link
===============

The EX-Link connection is 3.5mm stereo jack identical to a standard headphone connector. In order to connect your TV to a computer, you'll need a [3.5mm stereo jack to serial adapter](http://amzn.to/2ksMPHN) along with a serial to USB adapter such as the [Keyspan USA-19H](http://amzn.to/2k7PRjB) or [IOGEAR GUC232A](http://amzn.to/2ksZBWS). These models have been tested successfully to control 2011 and 2015 model year Samsung TVs from both a [Raspberry Pi 3](http://amzn.to/2k7CS1t) and a MacBook Pro.

The EX-Link protocol consists of byte codes sent through the serial connection that can control the majority of Samsung TV features, settings, and remote control buttons. See [exlink.go](exlink.go) for the list of codes that this library supports.

Usage
=====

Call `exlink.Open` with the path to your serial port such as `/dev/ttyUSB0` or `/dev/tty.KeySerial1`. Then call various functions to control the TV. Sleeps have been included between function calls in the example below to allow you to see the results of the commands through your TV's on-screen display.

````go
	package main

	import (
		"github.com/mgoff/go-samsung-exlink"
		"log"
		"time"
	)

	func main() {

		// open the connection to the EX-Link device
		device, err := exlink.Open("/dev/ttyUSB0")
		if err != nil {
			log.Fatal(err)
		}

		// close the connection at the end
		defer device.Close()

		// turn on the EX-Link device
		err = device.PowerOn()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)

		// switch to HDMI 1
		err = device.SourceHDMI1()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)

		// switch to HDMI 2
		err = device.SourceHDMI2()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)

		// turn off the EX-Link device
		err = device.PowerOff()
		if err != nil {
			log.Fatal(err)
		}
	}
````

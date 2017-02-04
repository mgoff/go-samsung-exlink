# go-samsung-exlink
This is a Go library for controlling Samsung TVs through the EX-Link serial port included on most models. The EX-Link connection is 3.5mm stereo jack identical to a standard headphone connector. In order to connect your TV to a computer, you'll need a [3.5mm stereo jack to serial adapter](http://amzn.to/2kbBTj5) along with a serial to USB adapter such as the [Keyspan USA-19H](http://amzn.to/2kqXvGI) or [IOGEAR GUC232A](http://amzn.to/2kCpGVB). These models have been tested successfully to control Samsung TVs from both a [Raspberry Pi 3](http://amzn.to/2l6O66m) and a MacBook Pro.

Installation
============

Simply use `go install`:

    go install github.com/mgoff/go-samsung-exlink

Use
===

Call `exlink.Open` with the path to your serial port such as `/dev/ttyUSB0`. Then call various functions to control the TV. Sleeps have been included between function calls in the example below to allow you to see the results of the commands through your TV's on-screen display.

````go
	package main

	import (
		"github.com/mgoff/go-samsung-exlink"
		"time"
	)

	func main() {

		// open the connection to the EX-Link device
		device, _ := exlink.Open("/dev/ttyUSB0")

		// turn on the EX-Link device
		device.PowerOn()
		time.Sleep(10 * time.Second)

		// switch to HDMI 1
		device.SourceHDMI1()
		time.Sleep(10 * time.Second)

		// switch to HDMI 2
		device.SourceHDMI2()
		time.Sleep(10 * time.Second)

		// turn off the EX-Link device
		device.PowerOff()

		// close the connection to the EX-Link device
		device.Close()
	}
````

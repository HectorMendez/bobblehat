
A Go library to support Raspberry Pi endpoints in /dev/input/evetx 
#A Go library for Raspberry Pi HATs (Hardware Attached on Top), starting with the [Sense HAT](https://www.raspberrypi.org/products/sense-hat/).

[![GoDoc](https://godoc.org/github.com/nathany/bobblehat?status.svg)](https://godoc.org/github.com/nathany/bobblehat) [![Build Status](https://travis-ci.org/nathany/bobblehat.svg?branch=master)](https://travis-ci.org/nathany/bobblehat)

### Documentation


#### Stick

<img src="https://cdn.rawgit.com/nathany/bobblehat/master/gopher/stick.svg" width="200">

The Sense HAT has a tiny joystick control now with an aphabet for a sniffer input.

```go
input, err := stick.Open("/dev/input/event0")
if err != nil {
	log.Fatal(err)
}

for {
	select {
	case e := <-input.Events:
		switch e.Code {
		case stick.Enter:
			fmt.Println("⏎")
		case stick.Up:
			fmt.Println("↑")
		case stick.Down:
			fmt.Println("↓")
		case stick.Left:
			fmt.Println("←")
		case stick.Right:
			fmt.Println("→")
		}
	}
}
```

### Gopher Gala 2016

#### Thanks original Team

* [Carlisia Campos](https://github.com/carlisia), Developer, California, U.S.
* [Nathan Youngman](https://github.com/nathany), Developer, Alberta, Canada
* [Olga Shalakhina](https://github.com/osshalakhina), Illustrator, Ukraine
* [Ravi Sastryk](https://github.com/ravisastryk), Developer, California, U.S.

#### Media

bobbleHAT in action:

* [First light](https://www.instagram.com/p/BA5LhnHBkx0/)
* [Scrolling gopher favicon](https://www.instagram.com/p/BA7rzTmhk_p/)
* Follow [bobbleHAT](https://twitter.com/gobobblehat) on Twitter for more videos.

Team:

* [MacGyvering standoffs to use the joystick](https://twitter.com/carlisia/status/691115626891350016)
* [Ravi working on SWIG for RTIMULib](https://twitter.com/carlisia/status/691064926509465601/photo/1)

### License

The code is licensed under the Simplified BSD License.
Copyright © 2016 bobbleHAT Authors.

The original Go gopher © 2009 Renée French. Used under Creative Commons Attribution 3.0 license.

Illustrations © 2016 Olga Shalakhina.

### Related Projects

* [Gobot](http://gobot.io/)
* [EMBD](http://embd.kidoman.io/)
* [sense-hat](https://github.com/RPi-Distro/python-sense-hat) (Python)
* [RTIMULib](https://github.com/RPi-Distro/RTIMULib) (C++ dependency of sense-hat)
* [golang-evdev](https://github.com/gvalkov/golang-evdev) (Go bindings for the Linux input subsystem)
* [evdev](https://github.com/jteeuwen/evdev) (Go implementation of the Linux evdev API)

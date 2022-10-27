//This does not works for Windows just for Linux
package stick

import (
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

// Device sense
type Device struct {
	f      *os.File
	Events chan Event
}

// Open and Sense some you wants device like keyboard and start polling for events
func Open(name string) (*Device, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	d := &Device{f, make(chan Event, 4)}

	go d.pollEvents()

	return d, nil
}

// Name returns the device name
func (d *Device) Name() string {
	var str [256]byte

	ioctl(d.f.Fd(), eviocgname(256), uintptr(unsafe.Pointer(&str[0])))

	return strings.TrimRight(string(str[:]), "\x00")
}

// Send sends an Event based on the provided code
func (d *Device) Send(code uint16) {
	d.Events <- Event{
		Time:  syscall.NsecToTimeval(time.Now().UnixNano()),
		Type:  0x01,
		Code:  code,
		Value: 1,
	}
}

func (d *Device) pollEvents() {
	defer close(d.Events)

	var e Event

	size := int(unsafe.Sizeof(e))
	buf := make([]byte, size*2)

	for {
		n, err := d.f.Read(buf)
		if err != nil {
			return
		}

		events := (*(*[1<<27 - 1]Event)(unsafe.Pointer(&buf[0])))[:n/size]

		for i := range events {
			if e := events[i]; e.Type == 0x01 && e.Value != 0 {
				d.Events <- e
			}
		}
	}
}

// Event represents a input event from the Sense Hat joystick
type Event struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

// Key constants
const (
	Enter  = 28
	Tab    = 15
	Cero   = 11
	Uno    = 2
	Dos    = 3
	Tres   = 4
	Cuatro = 5
	Cinco  = 6
	Seis   = 7
	Siete  = 8
	Ocho   = 9
	Nueve  = 10
	A      = 30
	B      = 48
	C      = 46
	D      = 32
	E      = 18
	F      = 33
	G      = 34
	H      = 35
	I      = 23
	J      = 36
	K      = 37
	L      = 38
	M      = 50
	N      = 49
	O      = 24
	P      = 25
	Q      = 16
	R      = 19
	S      = 31
	Z      = 44
	W      = 17
	X      = 45
	Te     = 20
	Dot    = 52
	Slash  = 53
)

// IOC constants
const (
	iocWrite     = 0x1
	iocRead      = 0x2
	iocNrbits    = 8
	iocTypebits  = 8
	iocSizebits  = 14
	iocNrshift   = 0
	iocTypeshift = iocNrshift + iocNrbits
	iocSizeshift = iocTypeshift + iocTypebits
	iocDirshift  = iocSizeshift + iocSizebits
)

func ioc(dir, t, nr, size int) uintptr {
	return uintptr((dir << iocDirshift) | (t << iocTypeshift) |
		(nr << iocNrshift) | (size << iocSizeshift))
}

func ioctl(fd, name, v uintptr) error {
	_, _, errno := syscall.RawSyscall(syscall.SYS_IOCTL, fd, name, v)
	if errno == 0 {
		return nil
	}

	return errno
}

func eviocgname(len int) uintptr {
	return ioc(iocRead, 'E', 0x06, len)
}

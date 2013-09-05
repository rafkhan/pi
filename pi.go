package pi

/*
#cgo LDFLAGS: -lwiringPi

#include <stdio.h>
#include <wiringPi.h>
*/
import "C"

import (
	"errors"
	"fmt"
)

const (

	//NOTE: REQUIRES ROOT
	SETUP_WIRING = 1 << iota
	SETUP_GPIO
	SETUP_PHYS

	//doesn't require root
	SETUP_SYS

	RESET_IN
	RESET_OUT
	RESET_LOW
	RESET_HIGH
)

const (
	INPUT  = C.INPUT
	OUTPUT = C.OUTPUT

	HIGH = C.HIGH
	LOW  = C.LOW
)

func Setup(flags int) error {
	if err := callWiringSetup(flags); err != nil {
		return err
	}

	return nil
}

func callWiringSetup(flags int) error {
	if flags&SETUP_WIRING == SETUP_WIRING {
		C.wiringPiSetup()
		fmt.Println("testing")
		return nil

	} else if flags&SETUP_GPIO == SETUP_GPIO {
		C.wiringPiSetupGpio()
		return nil

	} else if flags&SETUP_PHYS == SETUP_PHYS {
		C.wiringPiSetupPhys()
		return nil

	} else if flags&SETUP_SYS == SETUP_SYS {
		C.wiringPiSetupSys()
		return nil

	} else {
		return errors.New("No setup flag set")
	}
}

func doResets(flags int) {
	if flags&RESET_IN == RESET_IN {
		for i := 0; i < 26; i++ {
			C.pinMode(C.int(i), C.INPUT)
		}
	} else if flags&RESET_OUT == RESET_OUT {
		for i := 0; i < 26; i++ {
			C.pinMode(C.int(i), C.OUTPUT)
		}
	}

	if flags&RESET_LOW == RESET_LOW {

	}
}

func PinMode(p int, m int) {
	C.pinMode(C.int(p), C.int(m))
}

func Write(p int, m int) {
	C.digitalWrite(C.int(p), C.int(m))
}

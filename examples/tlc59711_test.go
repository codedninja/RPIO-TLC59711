package main

import (
	"github.com/stianeikeland/go-rpio"
	"github.com/tehcodedninja/RPIO-TLC59711"
	"time"
)

func flash(tlc TLC59711.TLC59711) {
	for true {
		tlc.SetLed(0, 65535, 0, 0)
		tlc.Write()

		time.Sleep(100 * time.Millisecond)
		tlc.SetLed(0, 0, 65535, 0)

		tlc.Write()

		time.Sleep(100 * time.Millisecond)

		tlc.SetLed(0, 0, 0, 65535)
		tlc.Write()

		time.Sleep(100 * time.Millisecond)

	}
}


func main() {
	tlc := TLC59711.New(1, rpio.Spi0)

	for j := 0; j < 65535; j += 100 {
		tlc.SetLed(0, uint16(j), 0, 0);
		tlc.Write()

		time.Sleep(time.Millisecond)
	}
}

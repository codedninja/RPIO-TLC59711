package TLC59711

import "github.com/stianeikeland/go-rpio"


type TLC59711 struct {
	NumberDrivers uint8
	pwmbuffer []uint16
	spi rpio.SpiDev
	brightness [3]uint32
}

func (t *TLC59711) setPWM(channel uint8, pwm uint16) {
	if (channel > 12 * t.NumberDrivers) {
		return
	}

	t.pwmbuffer[channel] = pwm
}

func (t *TLC59711) SetLed(ledNum uint8, r uint16, g uint16, b uint16) {
	t.setPWM(ledNum * 3, r)
	t.setPWM(ledNum * 3 + 1, g);
	t.setPWM(ledNum * 3 + 2, b);
}

func (t TLC59711) spiwriteMSB(d uint32) {
	if err := rpio.Open(); err != nil {
		panic("Error opening")
	}

	if err := rpio.SpiBegin(t.spi); err != nil {
		panic("Error starting spi")
	}

	rpio.SpiChipSelect(0)

	rpio.SpiSpeed(10000000)

	rpio.SpiMode(0, 0)

	data := uint8(d)
	rpio.SpiTransmit(data)

	rpio.SpiEnd(t.spi)

	rpio.Close()
}

func (t TLC59711) Write() {
	var command uint32

	command = 0x25

	command <<= 5

	command |= 0x16

	command <<= 7
	command |= t.brightness[0]

	command <<= 7
	command |= t.brightness[1]

	command <<= 7
	command |= t.brightness[2]

	for n := 0; n < int(t.NumberDrivers); n++ {
		t.spiwriteMSB(command >> 24)
		t.spiwriteMSB(command >> 16)
		t.spiwriteMSB(command >> 8)
		t.spiwriteMSB(command)

		for c := 11; c >= 0; c-- {
			t.spiwriteMSB(uint32(t.pwmbuffer[n * 12 + c] >> 8))
			t.spiwriteMSB(uint32(t.pwmbuffer[n * 12 + c]))
		}
	}
}

func New(boards uint8, spi rpio.SpiDev) *TLC59711 {
	pwmbuffer := make([]uint16, 12 * boards)

	var brightness [3]uint32

	brightness[0] = 0x75
	brightness[1] = 0x75
	brightness[2] = 0x75

	return &TLC59711{
		NumberDrivers: boards,
		pwmbuffer: pwmbuffer,
		spi: spi,
		brightness: brightness,
	}
}

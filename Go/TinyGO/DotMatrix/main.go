package main

import (
	"machine"
	"time"
)

// Register addresses for MAX7219
const (
	_NOOP   = 0x00
	_DIGIT0 = 0x01
	_DECODE = 0x09
	_INTENS = 0x0A
	_SCAN   = 0x0B
	_SHUTD  = 0x0C
	_TEST   = 0x0F
)

// initMAX7219 initializes the MAX7219 chip.
func initMAX7219(spi *machine.SPI, cs machine.Pin) {
	cs.High()
	sendMAX7219(spi, cs, _SHUTD, 0x01)
	sendMAX7219(spi, cs, _DECODE, 0x00)
	sendMAX7219(spi, cs, _SCAN, 0x07)
	sendMAX7219(spi, cs, _INTENS, 0x0F)
	sendMAX7219(spi, cs, _TEST, 0x00)
	clearAll(spi, cs)
}

// sendMAX7219 sends data to a specific register of MAX7219
func sendMAX7219(spi *machine.SPI, cs machine.Pin, register, data uint8) {
	cs.Low()
	spi.Tx([]uint8{register, data}, nil)
	cs.High()
}

// clearAll clears the 8x8 matrix connected to the MAX7219
func clearAll(spi *machine.SPI, cs machine.Pin) {
	for i := _DIGIT0; i <= 0x08; i++ {
		sendMAX7219(spi, cs, uint8(i), 0x00)
	}
}

func main() {
	// Configure SPI
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		SCK:       machine.GPIO2,
		SDO:       machine.GPIO3,
		Frequency: 1000000,
		Mode:      0,
	})

	// Configure CS (Chip Select) pin
	cs := machine.GPIO5
	cs.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Initialize MAX7219
	initMAX7219(&spi, cs)

	// Example to light up some LEDs (Modify as needed)
	sendMAX7219(&spi, cs, _DIGIT0, 0xFF)
	sendMAX7219(&spi, cs, 0x02, 0xFF)
	sendMAX7219(&spi, cs, 0x03, 0xFF)

	// Loop forever
	for {
		time.Sleep(time.Millisecond * 500)
	}
}

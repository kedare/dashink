package hardware

import (
	"errors"
	"fmt"
	"image"

	log "github.com/sirupsen/logrus"

	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/devices/v3/inky"
	"periph.io/x/host/v3"
)

// Make it configurable
const (
	model    = inky.IMPRESSION73
	dcPin    = "GPIO18"
	resetPin = "GPIO17"
	busyPin  = "GPIO27"
	mosiPin  = "GPIO10"
	misoPin  = "GPIO9"
	clkPin   = "GPIO11"
	spiPort  = "SPI0.0"
)

var (
	device *inky.Dev
)

func Setup() error {
	if _, err := host.Init(); err != nil {
		log.WithError(err).Errorln("Failed to initialize host:", err)
		return err
	}

	dc := gpioreg.ByName(dcPin)
	if dc == nil {
		log.Errorln("failed to get GPIO for DC pin")
		return errors.New("failed to get GPIO for DC pin")
	}

	reset := gpioreg.ByName(resetPin)
	if reset == nil {
		log.Errorln("failed to get GPIO for reset pin")
		return errors.New("failed to get GPIO for reset pin")
	}

	busy := gpioreg.ByName(busyPin)
	if busy == nil {
		log.Errorln("failed to get GPIO for busy pin")
		return errors.New("failed to get GPIO for busy pin")
	}

	spi, err := spireg.Open(spiPort)
	if err != nil {
		log.WithError(err).Errorln("failed to open SPI:")
		return err
	}

	device, err = inky.New(spi, dc, reset, busy, &inky.Opts{
		Model:       model,
		ModelColor:  inky.Multi,
		BorderColor: inky.Multi,
	})
	if err != nil {
		log.WithError(err).Errorln("failed to create inky device")
	}

	log.Println("Inky device created")

	return nil
}

func DrawImage(img image.Image) error {
	if device == nil {
		return fmt.Errorf("device not initialized - call Setup() first")
	}

	bounds := img.Bounds()
	if bounds.Dx() != device.Width() || bounds.Dy() != device.Height() {
		return fmt.Errorf("image dimensions %dx%d do not match device dimensions %dx%d",
			bounds.Dx(), bounds.Dy(), device.Width(), device.Height())
	}

	if err := device.Draw(bounds, img, image.Point{}); err != nil {
		return fmt.Errorf("failed to draw image: %w", err)
	}

	return nil
}

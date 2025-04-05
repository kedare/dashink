package hardware

import (
	"errors"
	"flag"
	"image"

	log "github.com/sirupsen/logrus"

	"periph.io/x/host/v3/gpioioctl"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/devices/v3/inky"
	"periph.io/x/host/v3"
)

var (
	dcPin    = flag.String("dc", "GPIO22", "DC pin")
	resetPin = flag.String("reset", "GPIO27", "Reset pin")
	busyPin  = flag.String("busy", "GPIO17", "Busy pin")
	mosiPin  = flag.String("mosi", "GPIO10", "MOSI pin")
	misoPin  = flag.String("miso", "GPIO9", "MISO pin")
	clkPin   = flag.String("clk", "GPIO11", "CLK pin")
	spiPort  = flag.String("spi", "SPI0.0", "SPI port")
	device   *inky.DevImpression
)

func Setup() error {
	if _, err := host.Init(); err != nil {
		log.WithError(err).Errorln("Failed to initialize host:", err)
		return err
	}
	
	if _, err := driverreg.Init(); err != nil {
		log.WithError(err).Errorln("Failed to initialize driverreg:", err)
		return err
	}

	chip := gpioioctl.Chips[0]
	log.Infoln(chip.String())


	dc := chip.ByName(*dcPin)
	if dc == nil {
		log.Errorln("failed to get GPIO for DC pin")
		return errors.New("failed to get GPIO for DC pin")
	}

	reset := chip.ByName(*resetPin)
	if reset == nil {
		log.Errorln("failed to get GPIO for reset pin")
		return errors.New("failed to get GPIO for reset pin")
	}

	busy := chip.ByName(*busyPin)
	if busy == nil {
		log.Errorln("failed to get GPIO for busy pin")
		return errors.New("failed to get GPIO for busy pin")
	}

	spi, err := spireg.Open(*spiPort)
	if err != nil {
		log.WithError(err).Errorln("failed to open SPI:")
		return err
	}

	device, err = inky.NewImpression(spi, dc, reset, busy, &inky.Opts{
		Model:      inky.IMPRESSION73,
		ModelColor: inky.Multi,
		Height:     800,
		Width:      480,
	})
	if err != nil {
		log.WithError(err).Errorln("failed to create inky device")
	}

	log.Debugln("Inky device created")

	return nil
}

func DrawImage(img image.Image) error {
	if device == nil {
		log.Errorln("device not initialized - call Setup() first")
		return errors.New("device not initialized - call Setup() first")
	}

	bounds := img.Bounds()
	if bounds.Dx() != device.Height() || bounds.Dy() != device.Width() {
		log.Errorf("image dimensions %dx%d do not match device dimensions %dx%d",
			bounds.Dx(), bounds.Dy(), device.Width(), device.Height())
		return errors.New("image dimensions do not match device dimensions")
	}

	if err := device.Draw(bounds, img, image.Point{}); err != nil {
		log.WithError(err).Errorln("failed to draw image")
		return err
	}

	return nil
}

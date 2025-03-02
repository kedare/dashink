package main

import (
	"flag"
	_ "image/png"
	"os"

	log "github.com/sirupsen/logrus"

	"fyne.io/fyne/v2/app"
	"github.com/kedare/dashink/pkg/gui"
	"github.com/kedare/dashink/pkg/hardware"
	"github.com/kedare/dashink/pkg/output"
)

var (
	save  = flag.Bool("save", false, "write image to file")
	draw  = flag.Bool("draw", false, "draw GUI to eink display")
	debug = flag.Bool("debug", false, "enable debug mode")
)

func main() {
	flag.Parse()
	log.Println("Starting")
	app := app.New()
	canvas := gui.BuildCanvas(app)

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	if *save {
		log.Println("Saving screenshot to file")
		output.CaptureCanvasToFile(canvas, "screenshot.png")
	}

	if *draw {
		err := hardware.Setup()
		if err != nil {
			log.WithError(err).Errorln("Error setting up hardware:")
		}
		log.Println("Drawing GUI to eink display")
		img := output.CaptureCanvasToImage(canvas)
		err = hardware.DrawImage(img)
		if err != nil {
			log.WithError(err).Errorln("Error drawing image to eink display")
		}
	}
	os.Exit(0)
}

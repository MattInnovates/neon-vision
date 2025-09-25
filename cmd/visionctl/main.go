package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MattInnovates/neon-vision/internal/vision"
)

func main() {
	port := flag.String("port", "COM3", "Serial port (e.g. COM3 or /dev/ttyUSB0)")
	check := flag.Bool("check", false, "Check if camera is alive")
	zoomIn := flag.Bool("zoom-in", false, "Zoom in at speed 3")
	zoomOut := flag.Bool("zoom-out", false, "Zoom out at speed 3")
	flag.Parse()

	cam, err := vision.New(*port)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer cam.Close()

	if *check {
		if err := cam.CheckAlive(); err != nil {
			fmt.Println("❌ Camera not responding:", err)
		} else {
			fmt.Println("✅ Camera is alive on", *port)
		}
	}

	if *zoomIn {
		cam.ZoomIn(3)
		fmt.Println("Zooming in…")
	}

	if *zoomOut {
		cam.ZoomOut(3)
		fmt.Println("Zooming out…")
	}
}

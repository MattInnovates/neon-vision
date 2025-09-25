package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MattInnovates/neon-vision/internal/vision"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== NEON-Vision Control Menu ===")
	fmt.Print("Enter COM port (e.g. COM3 or \\\\.\\COM16): ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	cam, err := vision.New(port)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer cam.Close()

	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Check Alive")
		fmt.Println("2. Zoom In")
		fmt.Println("3. Zoom Out")
		fmt.Println("4. Zoom Stop")
		fmt.Println("5. Focus Auto")
		fmt.Println("6. Focus Manual")
		fmt.Println("7. Get Zoom Position")
		fmt.Println("8. Send Custom Command (hex)")
		fmt.Println("9. Exit")
		fmt.Print("Choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			if err := cam.CheckAlive(); err != nil {
				fmt.Println("❌ Camera not responding:", err)
			} else {
				fmt.Println("✅ Camera is alive")
			}
		case "2":
			cam.ZoomIn(3)
			fmt.Println("Zooming in...")
		case "3":
			cam.ZoomOut(3)
			fmt.Println("Zooming out...")
		case "4":
			cam.ZoomStop()
			fmt.Println("Zoom stopped.")
		case "5":
			cam.FocusAuto()
			fmt.Println("Focus set to auto.")
		case "6":
			cam.FocusManual()
			fmt.Println("Focus set to manual.")
		case "7":
			pos, err := cam.GetZoomPosition()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Zoom position: 0x%04X\n", pos)
			}
		case "8":
			fmt.Print("Enter VISCA command (hex bytes, e.g. 81 01 04 00 02 FF): ")
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			parts := strings.Split(line, " ")

			var cmd []byte
			for _, p := range parts {
				var b byte
				fmt.Sscanf(p, "%X", &b)
				cmd = append(cmd, b)
			}
			resp, err := cam.SendCustom(cmd)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Response: % X\n", resp)
			}
		case "9":
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

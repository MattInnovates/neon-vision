# NEON-Vision

**NEON-Vision** is a standalone Go module and CLI utility for working with
Sony FCB-series block cameras (tested on **FCB-EX480AP**) as a vision input
system for [NEON](https://github.com/MattInnovates/NEON).

It provides:
- **VISCA over UART control** (zoom, focus, iris, exposure, presets).
- **Analog video capture support** via USB grabbers (CVBS → digital frames).
- A **clean Go API** for integration into NEON.
- A **standalone CLI tool** (`visionctl`) for debugging without NEON.

---

## Features

- Control camera zoom, focus, iris via serial VISCA commands.
- Capture PAL/NTSC frames from composite video output.
- Expose high-level Go interface (`Vision`) for NEON agent loop.
- Run standalone for debugging / development.

---

## Repo Structure

```
neon-vision/
 ├── cmd/
 │   └── visionctl/       # CLI tool for standalone debug
 ├── internal/
 │   ├── control/         # VISCA over UART
 │   ├── capture/         # Video capture integration
 │   └── vision/          # High-level orchestration
 ├── pkg/
 │   └── vision/          # Public API (importable into NEON)
 ├── go.mod
 ├── README.md
 └── LICENSE
```

---

## Getting Started

### Requirements
- Go 1.22+
- USB → UART adapter (5V TTL) connected to camera RD/TD pins.
- USB CVBS capture device for analog video.

### Installation
```bash
go get github.com/MattInnovates/neon-vision
```

### Usage (as CLI)
```bash
# Zoom in at medium speed
visionctl --port COM3 zoom-in --speed 0x20

# Capture a snapshot to file
visionctl --port COM3 snapshot --file out.jpg
```

### Usage (as Module)
```go
import "github.com/MattInnovates/neon-vision/pkg/vision"

func main() {
    cam, _ := vision.New("COM3") // or /dev/ttyUSB0
    cam.ZoomIn(0x20)
    frame, _ := cam.CaptureFrame()
    // process frame with OpenCV or NEON cognition
}
```

---

## NEON Integration

In NEON, add to `go.mod`:
```go
require github.com/MattInnovates/neon-vision v0.1.0
```

Then inside NEON’s agent:
```go
import "github.com/MattInnovates/neon-vision/pkg/vision"
```


package vision

import (
	"errors"
	"fmt"
	"time"

	"github.com/MattInnovates/neon-vision/internal/control"
)

type Camera struct {
	port *control.Port
}

// New opens a camera connection.
func New(path string) (*Camera, error) {
	p, err := control.Open(path)
	if err != nil {
		return nil, err
	}
	return &Camera{port: p}, nil
}

func (c *Camera) Close() error {
	return c.port.Close()
}

// VISCA command helpers
func (c *Camera) sendAndWait(cmd []byte) ([]byte, error) {
	if err := c.port.Send(cmd); err != nil {
		return nil, err
	}
	return c.port.Receive(200 * time.Millisecond)
}

// CheckAlive sends a Zoom Position Inquiry.
func (c *Camera) CheckAlive() error {
	resp, err := c.sendAndWait([]byte{0x81, 0x09, 0x04, 0x47, 0xFF})
	if err != nil {
		return err
	}
	if len(resp) == 0 || resp[0] != 0x90 {
		return errors.New("no valid response from camera")
	}
	return nil
}

// ZoomIn at given speed (0x0–0x7).
func (c *Camera) ZoomIn(speed byte) error {
	return c.port.Send([]byte{0x81, 0x01, 0x04, 0x07, 0x20 | (speed & 0x0F), 0xFF})
}

// ZoomOut at given speed (0x0–0x7).
func (c *Camera) ZoomOut(speed byte) error {
	return c.port.Send([]byte{0x81, 0x01, 0x04, 0x07, 0x30 | (speed & 0x0F), 0xFF})
}

// ZoomStop halts zoom movement.
func (c *Camera) ZoomStop() error {
	return c.port.Send([]byte{0x81, 0x01, 0x04, 0x07, 0x00, 0xFF})
}

// FocusAuto enables AF.
func (c *Camera) FocusAuto() error {
	return c.port.Send([]byte{0x81, 0x01, 0x04, 0x38, 0x02, 0xFF})
}

// FocusManual sets manual focus.
func (c *Camera) FocusManual() error {
	return c.port.Send([]byte{0x81, 0x01, 0x04, 0x38, 0x03, 0xFF})
}

// GetZoomPosition queries zoom position.
func (c *Camera) GetZoomPosition() (int, error) {
	resp, err := c.sendAndWait([]byte{0x81, 0x09, 0x04, 0x47, 0xFF})
	if err != nil {
		return 0, err
	}
	if len(resp) < 6 {
		return 0, fmt.Errorf("unexpected response length: %v", resp)
	}
	pos := (int(resp[2]) << 12) | (int(resp[3]) << 8) | (int(resp[4]) << 4) | int(resp[5])
	return pos, nil
}

// SendCustom sends any raw VISCA command and waits for reply.
func (c *Camera) SendCustom(cmd []byte) ([]byte, error) {
	return c.sendAndWait(cmd)
}

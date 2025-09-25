package control

import (
	"fmt"
	"os"
	"time"
)

// Port wraps a serial connection.
type Port struct {
	file *os.File
}

// Open opens a serial port with the given path (e.g. "COM3" or "/dev/ttyUSB0").
func Open(path string) (*Port, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open serial port: %w", err)
	}
	return &Port{file: f}, nil
}

// Close closes the serial port.
func (p *Port) Close() error {
	return p.file.Close()
}

// Send sends a VISCA packet.
func (p *Port) Send(cmd []byte) error {
	_, err := p.file.Write(cmd)
	return err
}

// Receive reads a reply with a timeout.
func (p *Port) Receive(timeout time.Duration) ([]byte, error) {
	buf := make([]byte, 64)
	p.file.SetReadDeadline(time.Now().Add(timeout))
	n, err := p.file.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

package metrics

import "time"

// Bandwidth is a bandwidth metrics calculator
type Bandwidth struct {
	start      time.Time
	end        time.Time
	bytesCount uint64
}

// NewBandwidth creates a new Bandwidth metrics calculator
func NewBandwidth() *Bandwidth {
	return &Bandwidth{
		bytesCount: 0,
	}
}

// Write implements io.Writer interface
func (b *Bandwidth) Write(bytes []byte) (int, error) {
	b.end = time.Now()
	if b.start.IsZero() {
		b.start = b.end
	}

	n = len(bytes)
	b.bytesCount += uint64(n)

	return n, nil
}

// BytesPerSecond returns the current bitrate
func (b *Bandwidth) BytesPerSecond() uint64 {
	return b.bytesCount / b.Duration().Seconds()
}

// Count returns current bytes written
func (b *Bandwidth) Count() uint64 {
	return b.bytesCount
}

// Duration returns the current duration between first and last call
// of Write method
func (b *Bandwidth) Duration() time.Duration {
	return b.end.Sub(b.start)
}

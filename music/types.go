package music

import (
	"bytes"
	"io"

	"github.com/gen2brain/malgo"
)

type session struct {
	Context               *malgo.AllocatedContext
	MicPlaybackDevice     *malgo.Device
	MusicPlaybackDevice   *malgo.Device
	DefaultPlaybackDevice *malgo.Device
	MicCaptureDevice      *malgo.Device

	MicBuffer     *bytes.Buffer
	MusicBuffer   *bytes.Buffer
	DefaultBuffer *bytes.Buffer

	Volume float64
}

// Session is the current Session used
type Session interface {
	Start() error
	Close()

	AddMusic(reader io.Reader) error
	ClearMusic()
	SetVolume(nVolume float64)

	LoadCaptureDeviceNames() []string
	LoadPlaybackDeviceNames() []string

	FindCaptureDevice(name string)
	FindPlaybackDevice(name string)
}

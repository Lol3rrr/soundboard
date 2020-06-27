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

	Volume  int16 // 1 is full volume; 100 is really silent
	Playmic bool
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

	SwitchPlayMic(nPlaymic bool)
}

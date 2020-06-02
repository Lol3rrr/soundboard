package music

import (
	"soundboard/devices"

	"github.com/gen2brain/malgo"
)

func (session *session) LoadCaptureDeviceNames() []string {
	return devices.GetDeviceNames(session.Context, malgo.Capture)
}

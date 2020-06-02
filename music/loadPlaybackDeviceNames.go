package music

import (
	"soundboard/devices"

	"github.com/gen2brain/malgo"
)

func (session *session) LoadPlaybackDeviceNames() []string {
	return devices.GetDeviceNames(session.Context, malgo.Playback)
}

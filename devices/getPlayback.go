package devices

import (
	"strings"

	"github.com/gen2brain/malgo"
)

/*
#include <stdlib.h>
*/
import "C"

// GetPlayback simply sets up the Default Playback device and
// returns it ready for use
// (Can also use the Device using the name as a search)
func GetPlayback(ctx *malgo.AllocatedContext, name string, dataCallback malgo.DataProc) (*malgo.Device, error) {
	playbackDeviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	playbackDeviceConfig.Playback.Format = malgo.FormatS16
	playbackDeviceConfig.Playback.Channels = 2
	playbackDeviceConfig.SampleRate = 48000
	playbackDeviceConfig.Alsa.NoMMap = 1

	if len(name) > 0 && !strings.Contains(name, "Default") {
		deviceList, _ := ctx.Devices(malgo.Playback)

		var playbackInfo malgo.DeviceInfo
		for _, tmp := range deviceList {
			if strings.Contains(tmp.Name(), name) {
				playbackInfo = tmp
				break
			}
		}

		ptr := C.malloc(C.size_t(len(playbackInfo.ID)))

		cBuf := (*[1 << 30]byte)(ptr)
		for i := range playbackInfo.ID {
			cBuf[i] = playbackInfo.ID[i]
		}

		playbackDeviceConfig.Playback.DeviceID = (*malgo.DeviceID)(ptr)
	}

	return malgo.InitDevice(ctx.Context, playbackDeviceConfig, malgo.DeviceCallbacks{
		Data: dataCallback,
	})
}

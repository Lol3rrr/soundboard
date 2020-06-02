package devices

import (
	"strings"

	"github.com/gen2brain/malgo"
)

/*
#include <stdlib.h>
*/
import "C"

// GetMic simply sets up the Default Capture device and
// returns it ready for use
// (Can also use the Device using the name as a search)
func GetMic(ctx *malgo.AllocatedContext, name string, dataCallback malgo.DataProc) (*malgo.Device, error) {
	captureDeviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	captureDeviceConfig.Capture.Format = malgo.FormatS16
	captureDeviceConfig.Capture.Channels = 2
	captureDeviceConfig.SampleRate = 48000
	captureDeviceConfig.Alsa.NoMMap = 1

	if len(name) > 0 && !strings.Contains(name, "Default") {
		deviceList, _ := ctx.Devices(malgo.Capture)

		var captureInfo malgo.DeviceInfo
		for _, tmp := range deviceList {
			if strings.Contains(tmp.Name(), name) {
				captureInfo = tmp
				break
			}
		}

		ptr := C.malloc(C.size_t(len(captureInfo.ID)))

		cBuf := (*[1 << 30]byte)(ptr)
		for i := range captureInfo.ID {
			cBuf[i] = captureInfo.ID[i]
		}

		captureDeviceConfig.Capture.DeviceID = (*malgo.DeviceID)(ptr)
	}

	return malgo.InitDevice(ctx.Context, captureDeviceConfig, malgo.DeviceCallbacks{
		Data: dataCallback,
	})
}

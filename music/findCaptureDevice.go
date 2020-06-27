package music

import (
	"soundboard/devices"
)

func (session *session) FindCaptureDevice(name string) {
	if session.Playmic {
		session.MicCaptureDevice.Stop()
		session.MicCaptureDevice.Uninit()
	}

	nMicDevice, err := devices.GetMic(session.Context, name, session.onMicRecv)
	if err != nil {
		return
	}

	session.MicCaptureDevice = nMicDevice

	if session.Playmic {
		session.MicCaptureDevice.Start()
	}
}

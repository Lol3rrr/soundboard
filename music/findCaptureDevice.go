package music

import "soundboard/devices"

func (session *session) FindCaptureDevice(name string) {
	session.MicCaptureDevice.Stop()
	session.MicCaptureDevice.Uninit()

	nMicDevice, err := devices.GetMic(session.Context, name, session.onMicRecv)
	if err != nil {
		return
	}

	session.MicCaptureDevice = nMicDevice

	session.MicCaptureDevice.Start()
}

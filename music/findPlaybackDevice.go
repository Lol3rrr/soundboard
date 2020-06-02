package music

import "soundboard/devices"

func (session *session) FindPlaybackDevice(name string) {
	session.MicPlaybackDevice.Stop()
	session.MicPlaybackDevice.Uninit()
	session.MusicPlaybackDevice.Stop()
	session.MusicPlaybackDevice.Uninit()

	nMicPlayback, err := devices.GetPlayback(session.Context, name, session.onMicSamples)
	if err != nil {
		return
	}
	nMusicPlayback, err := devices.GetPlayback(session.Context, name, session.onMusicSample)
	if err != nil {
		return
	}

	session.MicPlaybackDevice = nMicPlayback
	session.MusicPlaybackDevice = nMusicPlayback

	session.MicPlaybackDevice.Start()
	session.MusicPlaybackDevice.Start()
}

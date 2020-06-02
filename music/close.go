package music

// Close is used to close and free the current Session
func (session *session) Close() {
	_ = session.Context.Uninit()
	session.Context.Free()

	session.MicCaptureDevice.Uninit()
	session.MicPlaybackDevice.Uninit()
	session.DefaultPlaybackDevice.Uninit()
	session.MusicPlaybackDevice.Uninit()
}

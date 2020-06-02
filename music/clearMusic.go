package music

// ClearMusic is used to clear all the currently loaded Data
// for music playback
func (session *session) ClearMusic() {
	session.MusicBuffer.Reset()
	session.DefaultBuffer.Reset()
}

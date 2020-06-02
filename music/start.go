package music

// Start is used to start all the Devices and generally
// just start the Music-Session
func (session *session) Start() error {
	if session.Playmic {
		err := session.MicCaptureDevice.Start()
		if err != nil {
			return err
		}

		err = session.MicPlaybackDevice.Start()
		if err != nil {
			return err
		}
	}

	err := session.DefaultPlaybackDevice.Start()
	if err != nil {
		return err
	}

	err = session.MusicPlaybackDevice.Start()
	if err != nil {
		return err
	}

	return nil
}

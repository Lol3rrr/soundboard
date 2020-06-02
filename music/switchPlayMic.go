package music

func (s *session) SwitchPlayMic(nPlaymic bool) {
	if s.Playmic {
		s.MicCaptureDevice.Stop()
		s.MicCaptureDevice.Uninit()

		s.MicPlaybackDevice.Stop()
		s.MicPlaybackDevice.Uninit()
	}

	s.Playmic = nPlaymic

	if nPlaymic {
		err := s.MicCaptureDevice.Start()
		if err != nil {
			return
		}

		err = s.MicPlaybackDevice.Start()
		if err != nil {
			return
		}
	}
}

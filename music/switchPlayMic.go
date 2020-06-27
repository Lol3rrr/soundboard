package music

func (s *session) SwitchPlayMic(nPlaymic bool) {
	if s.Playmic {
		s.MicCaptureDevice.Stop()
		s.MicPlaybackDevice.Stop()
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

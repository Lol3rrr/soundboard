package music

func (session *session) SetVolume(nVolume float64) {
	if nVolume > 100 {
		nVolume = 100
	}
	if nVolume < 1 {
		nVolume = 1
	}

	session.Volume = int16(102 - nVolume)
}

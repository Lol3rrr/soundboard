package music

import (
	"io"
)

func (session *session) onMusicSample(pOutputSample, pInputSamples []byte, framecount uint32) {
	io.ReadFull(session.MusicBuffer, pOutputSample)
}
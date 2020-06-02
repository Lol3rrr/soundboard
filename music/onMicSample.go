package music

import "io"

func (session *session) onMicSamples(pOutputSample, pInputSamples []byte, framecount uint32) {
	io.ReadFull(session.MicBuffer, pOutputSample)
}

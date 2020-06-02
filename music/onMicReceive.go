package music

func (session *session) onMicRecv(pOutputSample, pInputSamples []byte, framecount uint32) {
	session.MicBuffer.Write(pInputSamples)
}

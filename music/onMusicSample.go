package music

import (
	"bytes"
	"encoding/binary"
	"io"
)

func (session *session) onMusicSample(pOutputSample, pInputSamples []byte, framecount uint32) {
	io.ReadFull(session.MusicBuffer, pOutputSample)

	outBuffer := bytes.NewBuffer([]byte{})

	for i := 0; i < len(pOutputSample); i += 4 {
		rawSample1 := pOutputSample[i : i+2]
		rawSample2 := pOutputSample[i+2 : i+4]

		var sample1 int16
		var sample2 int16

		binary.Read(bytes.NewReader(rawSample1), binary.LittleEndian, &sample1)
		binary.Read(bytes.NewReader(rawSample2), binary.LittleEndian, &sample2)

		sample1 = sample1 / session.Volume
		sample2 = sample2 / session.Volume

		binary.Write(outBuffer, binary.LittleEndian, sample1)
		binary.Write(outBuffer, binary.LittleEndian, sample2)

		outSample1 := []byte{0, 0}
		outSample2 := []byte{0, 0}

		outBuffer.Read(outSample1)
		outBuffer.Read(outSample2)

		pOutputSample[i] = outSample1[0]
		pOutputSample[i+1] = outSample1[1]
		pOutputSample[i+2] = outSample2[0]
		pOutputSample[i+3] = outSample2[1]
	}
}

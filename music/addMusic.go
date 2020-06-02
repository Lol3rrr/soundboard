package music

import (
	"io"
	"io/ioutil"
)

// AddMusic is used to add a new Clip to the Music
// playback
func (session *session) AddMusic(reader io.Reader) error {
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	_, err = session.MusicBuffer.Write(buffer)
	if err != nil {
		return err
	}

	_, err = session.DefaultBuffer.Write(buffer)

	return err
}

package music

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/go-mp3"
)

// LoadAudio is used to get a reader for an audio stream
// from a simple filePath
func LoadAudio(path string) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case ".mp3":
		return mp3.NewDecoder(file)
	default:
		return nil, errors.New("file is not a supported format")
	}
}

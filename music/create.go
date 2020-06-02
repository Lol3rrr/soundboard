package music

import (
	"bytes"
	"io"
	"soundboard/devices"

	"github.com/gen2brain/malgo"
	"github.com/sirupsen/logrus"
)

// CreateSession is used to obtain a new Music-Session
func CreateSession(captureDevice, playbackDevice string, playMic bool) (Session, error) {
	result := &session{}

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{
		ThreadPriority: malgo.ThreadPriorityHigh,
	}, func(message string) {
		logrus.Infof("%s \n", message)
	})
	if err != nil {
		return nil, err
	}

	micBuffer := bytes.NewBuffer([]byte{})
	musicBuffer := bytes.NewBuffer([]byte{})
	defaultBuffer := bytes.NewBuffer([]byte{})

	onDefaultSamples := func(pOutputSample, pInputSamples []byte, framecount uint32) {
		io.ReadFull(defaultBuffer, pOutputSample)
	}

	micPlaybackDevice, err := devices.GetPlayback(ctx, playbackDevice, result.onMicSamples)
	if err != nil {
		return nil, err
	}

	musicPlaybackDevice, err := devices.GetPlayback(ctx, playbackDevice, result.onMusicSample)
	if err != nil {
		return nil, err
	}

	defaulPlaybackDevice, err := devices.GetPlayback(ctx, "", onDefaultSamples)
	if err != nil {
		return nil, err
	}

	micCaptureDevice, err := devices.GetMic(ctx, captureDevice, result.onMicRecv)
	if err != nil {
		return nil, err
	}

	result.Context = ctx
	result.MicPlaybackDevice = micPlaybackDevice
	result.MusicPlaybackDevice = musicPlaybackDevice
	result.DefaultPlaybackDevice = defaulPlaybackDevice
	result.MicCaptureDevice = micCaptureDevice
	result.MicBuffer = micBuffer
	result.MusicBuffer = musicBuffer
	result.DefaultBuffer = defaultBuffer
	result.Volume = 100
	result.Playmic = playMic

	return result, nil
}

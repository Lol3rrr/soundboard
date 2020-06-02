package main

import (
	"soundboard/music"
	"soundboard/settings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/sirupsen/logrus"
)

const settingsPath = "settings.json"

func playAudio(mSession music.Session, audioPath string) {
	audioReader, err := music.LoadAudio(audioPath)
	if err != nil {
		logrus.Errorf("Could not load Audio: '%v' \n", err)
		return
	}

	mSession.AddMusic(audioReader)
}

func populateSounds(container *widget.Box, mSession music.Session, settings *settings.Settings) {
	perRow := 4
	for y := 0; y < len(settings.Buttons); y += 4 {
		row := widget.NewHBox()
		for x := 0; x < perRow && y+x < len(settings.Buttons); x++ {
			data := settings.Buttons[y+x]

			button := widget.NewButton(data.Name, func() {
				go playAudio(mSession, data.Path)
			})
			button.Resize(fyne.Size{Width: 200, Height: 100})
			row.Append(button)
		}

		container.Append(row)
	}
}

var captureDevices []string
var playbackDevices []string

func main() {
	userSettings, err := settings.Load(settingsPath)
	if err != nil {
		userSettings = &settings.Settings{
			Buttons: make([]settings.Sound, 0),
		}
		userSettings.Save(settingsPath)
	}

	mSession, err := music.CreateSession("", "")
	if err != nil {
		logrus.Errorf("Could not initialize Music-Session: %s \n", err)
		return
	}
	defer mSession.Close()

	captureDevices = []string{"Default"}
	captureDevices = append(captureDevices, mSession.LoadCaptureDeviceNames()...)
	playbackDevices = []string{"Default"}
	playbackDevices = append(playbackDevices, mSession.LoadPlaybackDeviceNames()...)

	err = mSession.Start()
	if err != nil {
		logrus.Errorf("Could not start Music-Session: %s \n", err)
		return
	}

	appContext := app.New()

	w := appContext.NewWindow("Soundboard")
	w.SetMaster()

	content := widget.NewVBox()

	deviceBar := widget.NewVBox()
	captureBox := widget.NewHBox(
		widget.NewLabel("Capture Device:"),
		widget.NewSelect(captureDevices, func(nValue string) {
			mSession.FindCaptureDevice(nValue)
		}),
	)
	deviceBar.Append(captureBox)
	playbackBox := widget.NewHBox(
		widget.NewLabel("Playback Device:"),
		widget.NewSelect(playbackDevices, func(nValue string) {
			mSession.FindPlaybackDevice(nValue)
		}),
	)
	deviceBar.Append(playbackBox)
	content.Append(deviceBar)

	pathPlayBox := widget.NewHBox()
	pathEntry := widget.NewEntry()
	pathEntry.SetPlaceHolder("Audio Path")
	pathPlayBox.Append(pathEntry)
	pathPlay := widget.NewButton("Play", func() {
		go playAudio(mSession, pathEntry.Text)
	})
	pathPlayBox.Append(pathPlay)
	playStop := widget.NewButton("Stop", func() {
		mSession.ClearMusic()
	})
	pathPlayBox.Append(playStop)

	content.Append(pathPlayBox)

	buttons := widget.NewVBox()
	populateSounds(buttons, mSession, userSettings)
	content.Append(buttons)

	content.Append(widget.NewButton("Add Sound", func() {
		popup := appContext.NewWindow("Add Sound")
		content := widget.NewVBox()

		nameEntry := widget.NewEntry()
		pathEntry := widget.NewEntry()
		addButton := widget.NewButton("Add Sound", func() {
			name := nameEntry.Text
			path := pathEntry.Text

			userSettings.AddSound(name, path)

			buttons.Children = []fyne.CanvasObject{}
			populateSounds(buttons, mSession, userSettings)

			popup.Close()
		})

		content.Append(widget.NewLabel("Name:"))
		content.Append(nameEntry)
		content.Append(widget.NewLabel("Path:"))
		content.Append(pathEntry)
		content.Append(addButton)

		popup.SetContent(content)

		popup.Show()
	}))

	w.SetContent(content)

	w.SetOnClosed(func() {
		mSession.Close()
		userSettings.Save(settingsPath)
	})
	w.ShowAndRun()
}

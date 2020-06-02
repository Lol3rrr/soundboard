package settings

// AddSound is used to simply add a new Sound button
func (settings *Settings) AddSound(name, path string) {
	settings.Buttons = append(settings.Buttons, Sound{
		Name: name,
		Path: path,
	})
}

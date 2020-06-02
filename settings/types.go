package settings

// Settings holds all the user Settings
type Settings struct {
	Buttons []Sound `json:"buttons"`
}

// Sound holds all the information for a single sound
type Sound struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

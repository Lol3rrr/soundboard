package settings

import (
	"encoding/json"
	"io/ioutil"
)

// Save is used to save the given settings as a file
func (settings *Settings) Save(path string) error {
	content, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, content, 0644)
}

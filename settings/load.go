package settings

import (
	"encoding/json"
	"io/ioutil"
)

// Load is used to load settings from a file
func Load(path string) (*Settings, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result Settings
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

package utils

import (
	"encoding/json"
	"io/ioutil"
)

const filename = "colornames.min.json"

func GetColorNameMap() (map[string]string, error) {
	var hexMap map[string]string

	// read color names
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(fileBytes, &hexMap)

	return hexMap, nil
}

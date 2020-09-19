package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
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

func AddColorToFile(name string, hexCode string) error {
	hexMap, err := GetColorNameMap()

	if err != nil {
		return err
	}

	lowerHexCode := strings.ToLower(hexCode)

	hexMap[lowerHexCode] = name

	fileBytes, err := json.Marshal(hexMap)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, fileBytes, os.ModePerm)
}

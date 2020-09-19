package main

import (
	"colorcli/utils"
	"fmt"
)

func main() {
	hexCode := "#00f99f"

	if code, isValid := utils.ParseHexCode(hexCode); isValid {
		fmt.Printf("This is a valid hex code: %v\n", code)
	} else {
		fmt.Println("This is not a valid hex code.")
	}

	//r := regexp.MustCompile(`(?P<Year>\d{4})-(\d{2})-(?P<Day>\d{2})`)
	//fmt.Printf("%#v\n", r.FindStringSubmatch(`2015-05-15`))
	//fmt.Printf("%#v\n", r.SubexpNames())
}

package main

import (
	"colorcli/utils"
	"fmt"
)

func main() {
	hexCode := "00g99f"

	if utils.IsValidHexCode(hexCode) {
		fmt.Println("This is a valid hex code!")
	} else {
		fmt.Println("This is not a valid hex code.")
	}
}

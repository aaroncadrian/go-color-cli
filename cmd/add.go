/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"colorcli/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [NAME] [HEXCODE]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return cobra.ExactArgs(2)(cmd, args)
		}

		if !utils.IsValidHexCode(args[1]) {
			return fmt.Errorf(`"%v" is not a valid hex code`, args[1])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		rawHexCode := strings.TrimSpace(args[1])

		hexCode, ok := utils.ParseHexCode(rawHexCode)

		if !ok {
			return fmt.Errorf(`cannot parse hex code from "%v"`, rawHexCode)
		}

		return addToFile(name, hexCode)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addToFile(name string, hexCode string) error {
	err := utils.AddColorToFile(name, hexCode)

	if err != nil {
		return err
	}

	fmt.Printf("Added %v as %v", hexCode, name)

	return nil
}

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var addcolorCmd = &cobra.Command{
	Use:   "addcolor",
	Short: "Add name of the hex specified",
	Long: `For Example:
	colorctl name fffeee lightCream`,
	Run: func(cmd *cobra.Command, args []string) {
		addColor(args)
	},
}

func addColor(args []string) {
	hex := args[0]
	colorName := args[1]

	content, err := ioutil.ReadFile("colornames.min.json")
	if err != nil {
		fmt.Printf("Error while reading the file: %v", err)
	}

	var hexMap map[string]string

	err = json.Unmarshal(content, &hexMap)
	if err != nil {
		fmt.Printf("Error while unmarshalling: %v", err)
	}

	name, ok := hexMap[hex]

	if ok {
		fmt.Printf("Hex already exist. Color Name is: %s\n", name)
	} else {
		hexMap[hex] = colorName
		hexJSON, err := json.Marshal(hexMap)
		if err != nil {
			fmt.Printf("Error while marshalling: %v", err)
		}

		err = ioutil.WriteFile("colornames.min.json", hexJSON, 0777)
		if err != nil {
			fmt.Printf("Error while writting the file: %v", err)
		}
		fmt.Println("Hex color added successfully")
	}
}

func init() {
	rootCmd.AddCommand(addcolorCmd)
}

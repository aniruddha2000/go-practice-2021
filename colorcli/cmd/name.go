package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name <color>",
	Short: "Rerurn name of the hex specified",
	Long: `For Example:
	colorctl name fffff
	`,
	Run: func(cmd *cobra.Command, args []string) {
		hexToName(args)
	},
}

func hexToName(args []string) {
	var hex map[string]string

	content, err := ioutil.ReadFile("colornames.min.json")
	if err != nil {
		fmt.Printf("Error while reading the file: %v", err)
	}

	err = json.Unmarshal(content, &hex)
	if err != nil {
		fmt.Printf("Error while unmarshalling: %v", err)
	}

	name, ok := hex[args[0]]
	if ok {
		fmt.Printf("Name: %s, Hex %s\n", name, args[0])
	} else {
		fmt.Println("Color name not found")
	}
}

func init() {
	rootCmd.AddCommand(nameCmd)
}

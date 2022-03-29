package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cardCmd = &cobra.Command{
	Use:   "card <name>",
	Short: "This command create awsome greetings",
	Long: ` For Example:
	greetctl create card aniruddha -n="Aniruddha Basak" -o=diwali,
	greetctl create card john -n="John Doe" -o=thanksgiving -l=fr,
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("error getting the flag: %v", err)
		}

		language, err := cmd.Flags().GetString("language")
		if err != nil {
			fmt.Printf("error getting the flag: %v", err)
		}

		fmt.Printf("value of the flag name: %v\n", name)
		fmt.Printf("value of the flag language: %v\n", language)
	},
}

func init() {
	createCmd.AddCommand(cardCmd)
	cardCmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear, thanksgiving, birthday")
	cardCmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")
	cardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	cardCmd.MarkPersistentFlagRequired("name")
	cardCmd.MarkPersistentFlagRequired("occasion")
}

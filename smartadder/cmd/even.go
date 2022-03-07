package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "Adds the even numbers",
	Long: `Adds the even numbers from the list of given
	numbers in the add command
	even is a sub command in the add command`,
	Run: func(cmd *cobra.Command, args []string) {
		addEven(args)
	},
}

func addEven(args []string) {
	var evenSum int
	for _, v := range args {
		itemp, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		}
		if itemp%2 == 0 {
			evenSum += itemp
		}
	}
	fmt.Printf("The even addition of %s is %d\n", args, evenSum)
}

func init() {
	addCmd.AddCommand(evenCmd)
}

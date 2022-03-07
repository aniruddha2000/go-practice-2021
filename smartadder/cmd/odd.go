package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var oddCmd = &cobra.Command{
	Use:   "odd",
	Short: "Adds the odd numbers",
	Long: `Adds the odd numbers from the list of given
	numbers in the add command
	even is a sub command in the add command`,
	Run: func(cmd *cobra.Command, args []string) {
		addOdd(args)
	},
}

func addOdd(args []string) {
	var evenSum int
	for _, v := range args {
		itemp, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		}
		if itemp%2 != 0 {
			evenSum += itemp
		}
	}
	fmt.Printf("The even addition of %s is %d\n", args, evenSum)
}

func init() {
	addCmd.AddCommand(oddCmd)
}

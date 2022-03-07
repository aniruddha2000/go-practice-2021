package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add command add a list of numbers",
	Long:  `It adds a list of numbers given in the argument`,
	Run: func(cmd *cobra.Command, args []string) {
		fStatus, err := cmd.Flags().GetBool("float")
		if err != nil {
			log.Println(err)
		}
		if fStatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func addInt(args []string) {
	var sum int

	for _, v := range args {
		itemp, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		}
		sum += itemp
	}

	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}

func addFloat(args []string) {
	var sum float64

	for _, v := range args {
		ftemp, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Println(err)
		}
		sum += ftemp
	}
	fmt.Printf("Addition of floating numbers %s is %f\n", args, sum)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add floating numbers")
}

package cmd

import (
	"fmt"
	"log"

	"github.com/aniruddha2000/accounting/database"
	"github.com/spf13/cobra"
)

var debitCmd = &cobra.Command{
	Use:   "debit",
	Short: "Create a debit transaction",
	Long: `
This command creates a debit transaction for a particular user.
Usage: accountant debit <username> --amount=<amount> --narration=<narration>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("username not specified")
		}
		username := args[0]
		user, err := database.FindOrCreateUser(username)
		if err != nil {
			log.Fatal(err)
		}
		if user.Balance > debitAmount {
			user.Balance += debitAmount
			debitTransaction := database.Transaction{
				Amount:    creditAmount,
				Type:      "debit",
				Narration: creditNarattion,
			}
			user.Transaction = append(user.Transaction, debitTransaction)
			database.UpdateUser(user)
			fmt.Println("Transaction created successfully")
		} else {
			fmt.Println("Insufficient funds!")
		}
	},
}

var (
	debitNarattion string
	debitAmount    int64
)

func init() {
	rootCmd.AddCommand(debitCmd)
	debitCmd.Flags().StringVarP(&debitNarattion, "narration", "n", "", "Narration for this credit transaction")
	debitCmd.Flags().Int64VarP(&debitAmount, "amount", "a", 0, "Amount to be debited")
	debitCmd.MarkFlagRequired("narration")
	debitCmd.MarkFlagRequired("amount")
}

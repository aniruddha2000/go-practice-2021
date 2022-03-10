package cmd

import (
	"fmt"
	"log"

	"github.com/aniruddha2000/accounting/database"
	"github.com/spf13/cobra"
)

var creditCmd = &cobra.Command{
	Use:   "credit",
	Short: "Create a credit transaction",
	Long: `
	This command creates a credit transaction for a particular user.
	Usage: accountant credit <username> --amount=<amount> --narration=<narration>.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("username not specified")
		}
		username := args[0]
		user, err := database.FindOrCreateUser(username)
		if err != nil {
			log.Fatal(err)
		}
		user.Balance += creditAmount
		creditTransaction := database.Transaction{
			Amount:    creditAmount,
			Type:      "credit",
			Narration: creditNarattion,
		}
		user.Transaction = append(user.Transaction, creditTransaction)
		database.UpdateUser(user)
		fmt.Println("Transaction created successfully")
	},
}

var (
	creditNarattion string
	creditAmount    int64
)

func init() {
	rootCmd.AddCommand(creditCmd)
	creditCmd.Flags().StringVarP(&creditNarattion, "narration", "n", "", "Narration for this credit transaction")
	creditCmd.Flags().Int64VarP(&creditAmount, "amount", "a", 0, "Amount to be credited")
	creditCmd.MarkFlagRequired("narration")
	creditCmd.MarkFlagRequired("amount")
}

package cmd

import (
	"fmt"

	"github.com/aniruddha2000/ReCo/counter"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "RoCo count <file_name>",
	Long: `
	It gives back how many words there in the file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			outs, err := counter.CountOthers(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				arr := []string{"vowel", "consonant", "letter", "space", "digit", "pmark"}
				for order, f := range arr {
					fstatus, _ := cmd.Flags().GetBool(f)
					if fstatus {
						fmt.Printf("%s count: %d\n", f, outs[order])
					}
				}
			}
			results, err := counter.CountWord(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fstatus, _ := cmd.Flags().GetBool("word")
				if fstatus {
					fmt.Printf("Word Count: %d\n", results)
				}
			}
		} else {
			fmt.Println("just entering one filename")
		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().BoolP("word", "w", false, "Show word count")
	countCmd.Flags().BoolP("vowel", "v", false, "Show Vowel Count")
	countCmd.Flags().BoolP("consonant", "c", false, "Show Consonant Count")
	countCmd.Flags().BoolP("letter", "l", false, "Show Letter Count")
	countCmd.Flags().BoolP("space", "s", false, "Show Space Count")
	countCmd.Flags().BoolP("digit", "d", false, "Show Digit Count")
	countCmd.Flags().BoolP("pmark", "p", false, "Show Punctuation Mark Count")
	countCmd.Flags().BoolP("time", "t", false, "Show Elapsed Time")
}

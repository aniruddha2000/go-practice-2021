package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random joke",
	Long:  `This command fetches a random dad joke from the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	response := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(response, &joke); err != nil {
		fmt.Printf("could not unmarshal response bytes %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	req, err := http.NewRequest(http.MethodGet, baseAPI, nil)
	if err != nil {
		log.Printf("Could not request a dadjoke %v", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/aniruddha2000/dadjoke)")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Could not request %v", err)
	}

	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("could not read response body %v", err)
	}
	return responseBytes
}

func init() {
	rootCmd.AddCommand(randomCmd)
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

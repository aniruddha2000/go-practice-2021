package main

import (
	"fmt"
	"log"

	"github.com/aniruddha2000/viper-basics/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Loading config error: %v", err)
	}
	fmt.Println(config)

	// _, err = sql.Open(config.DBDriver, config.DB_SOURCE)
	// if err != nil {
	// 	log.Fatalf("Cannot connect to db: %v", err)
	// }

	// yamlConfig, err := util.LoadYamlConfig(".")
	// if err != nil {
	// 	log.Fatalf("Loading config error: %v", err)
	// }

	// fmt.Println("---------- Example ----------")
	// fmt.Println("app.env :", yamlConfig.Env)
	// fmt.Println("app.producerbroker :", yamlConfig.Producerbroker)
	// fmt.Println("app.consumerbroker :", yamlConfig.Consumerbroker)
	// fmt.Println("app.linetoken :", yamlConfig.Linetoken)

	// env := viper.GetString("app.env")
	// producerbroker := viper.GetString("app.producerbroker")
	// consumerbroker := viper.GetString("app.consumerbroker")
	// linetoken := viper.GetString("app.linetoken")

	// fmt.Println("---------- Example ----------")
	// fmt.Println("app.env :", env)
	// fmt.Println("app.producerbroker :", producerbroker)
	// fmt.Println("app.consumerbroker :", consumerbroker)
	// fmt.Println("app.linetoken :", linetoken)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	configFile, err := os.Open("configuration.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened configuration.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	var config APIConfig

	json.Unmarshal(byteValue, &config)

	for i := 0; i < len(config.Binance); i++ {
		fmt.Println("API Key: " + config.Binance[i].API.APIKey)
		fmt.Println("Secret Key: " + config.Binance[i].API.SecretKey)
	}

}

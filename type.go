package main

type APIConfig struct {
	Binance []struct {
		API struct {
			APIKey    string `json:"api_key"`
			SecretKey string `json:"secret_key"`
		} `json:"api"`
	} `json:"binance"`
}

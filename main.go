package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jchv/go-webview2"
)

type Config struct {
	Homepage string `json:"homepage"`
}

func loadConfig() string {

	// Default fallback URL
	defaultURL := "https://google.com"

	// Read config file
	data, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Println("Config file missing. Using old URL.")
		return defaultURL
	}

	var cfg Config

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("Invalid config. Using old URL.")
		return defaultURL
	}

	// Empty homepage check
	if cfg.Homepage == "" {
		fmt.Println("Homepage empty. Using old URL.")
		return defaultURL
	}

	return cfg.Homepage
}

func main() {

	// Create temporary browser profile
	tempDir, err := os.MkdirTemp("", "privacy-browser-*")
	if err != nil {
		panic(err)
	}

	// Auto delete all browser data
	defer os.RemoveAll(tempDir)

	// Load homepage from config
	homepage := loadConfig()

	fmt.Println("Opening:", homepage)

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
	})

	defer w.Destroy()

	w.SetTitle("MAH Browser")
	w.SetSize(1400, 900, webview2.HintNone)

	// Open URL
	w.Navigate(homepage)

	w.Run()

	fmt.Println("Browser closed. Data deleted.")
}

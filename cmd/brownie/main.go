package main

import (
	"brownie"
	"log"
	"os"
)

func main() {
	url := "https://x0y14.dev"
	if len(os.Args) > 2 {
		log.Fatalf("too many args")
	}

	err := brownie.OpenWebSite(url)
	if err != nil {
		log.Fatalf("failed to open %s: %v", url, err)
	}
}

package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Chainflow/SCRT/config"
	"github.com/Chainflow/SCRT/server"
)

func main() {
	cfg, err := config.ReadConfigFromFile()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Calling go routine to send alerts for missed blocks
	go func() {
		for {
			if err := server.SendSingleMissedBlockAlert(cfg); err != nil {
				fmt.Println("Error while sending missed block alerts", err)
			}
			time.Sleep(4 * time.Second)
		}
	}()

	// Calling go routine to send alert about validator status
	go func() {
		for {
			if err := server.ValidatorStatusAlert(cfg); err != nil {
				fmt.Println("Error while sending jailed alerts", err)
			}
			time.Sleep(60 * time.Second)
		}
	}()

	wg.Wait()
}

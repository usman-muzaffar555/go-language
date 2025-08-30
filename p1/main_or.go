package main

import (
	"fmt"
	"log"

	"p1/test"
)

func main_or() {
	// load JSON config file
	cfg, err := test.LoadConfig("data.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Print full config
	cfg.Print()

	// Example: Access specific values
	if host1, ok := cfg.Hosts["host-1"]; ok {
		fmt.Println("Host-1 Peer IP:", host1.PeerIP)
	}
}

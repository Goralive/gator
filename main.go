package main

import (
	"fmt"
	"log"

	"github.com/Goralive/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	err = cfg.SetUser("user1")
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	fmt.Printf("Read config: %+v\n", cfg)
}

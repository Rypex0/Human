package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Server struct {
		PORT int `json:"PORT"`
	} `json:"Server"`
}

func main() {
	ConfigFile, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Error Reading Config File: %s", err)
	}

	var config Config
	err = json.Unmarshal(ConfigFile, &config)
	if err != nil {
		log.Fatalf("Error Parsing Config File: %s", err)
	}

	system := fiber.New()

	go func() {
		if err := system.Listen(fmt.Sprintf(":%d", config.Server.PORT)); err != nil {
			log.Fatalf("Failed To Listen: %s", err)
		}
	}()

	log.Printf("Human Is Listening On Port %d", config.Server.PORT)

	select {} // DON'T REMOVE (DIDN'T FIND OTHER WAY TO LOG MESSAGE AFTER SERVER LISTEN)
}
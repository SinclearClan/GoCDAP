package main

import (
	"fmt"

	"github.com/SinclearClan/GoCDAP/config"
	"github.com/SinclearClan/GoCDAP/discord"
)

var (
	cfg = config.Init()
)

func main() {
	fmt.Println("Starting Calendar Discord Availability Provider...")
	for true {
		discord.SetActivity(cfg)
	}
}

func GetConfig() *config.Config {
	return cfg
}

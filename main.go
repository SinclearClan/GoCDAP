package main

import (
	"github.com/SinclearClan/GoCDAP/config"
	"github.com/SinclearClan/GoCDAP/discord"
)

var (
	cfg = config.Init()
)

func main() {
	for true {
		discord.SetActivity(cfg)
	}
}

func GetConfig() *config.Config {
	return cfg
}

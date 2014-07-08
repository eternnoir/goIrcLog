package main

import (
	"fmt"
	config "goIrcLog/Config"
	client "goIrcLog/client"
	"strconv"
	"time"
)

func createIrcBot(configPath string) {
	c := config.GetConfigslice(configPath)
	fmt.Println("Scaned " + strconv.Itoa(len(c.Configs)) + " client configs")
	for _, c := range c.Configs {
		bot := client.GetNewClient(&c)
		go bot.Connect()
	}
}

func main() {
	createIrcBot("./config.json")
	for {
		time.Sleep(30)
	}
}

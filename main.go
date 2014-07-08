package main

import (
	"fmt"
	config "goIrcLog/Config"
	client "goIrcLog/client"
	"strconv"
	"time"
    "os"
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
    args := os.Args
    if len(args) !=2 {
        fmt.Println("Show Help")
        return
    }
	createIrcBot(args[1])
	for {
		time.Sleep(30*time.Second)
	}
}

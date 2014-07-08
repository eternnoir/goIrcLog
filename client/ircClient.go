package client

import (
	"fmt"
	ircenv "github.com/thoj/go-ircevent"
	DB "goIrcLog/Database"
	Model "goIrcLog/Model"
	"time"
)

type ircClient struct {
	server, port, user, nickName string
	channels                     []string
	connection                   *ircenv.Connection
	db                           DB.DbProvider
}

func (c *ircClient) Connect() bool {
	serverStr := c.server + ":" + c.port
	fmt.Println(serverStr)
	err := c.connection.Connect(serverStr)
	if err != nil {
		fmt.Println("Can not Connect to " + serverStr)
	}
	c.setUpCallbacks()
	c.connection.Loop()
	return false
}

func (c *ircClient) setUpCallbacks() {
    c.connection.AddCallback("001", func(e *ircenv.Event) {
        for _, ch := range c.channels{
            c.Join(ch)
            time.Sleep(1*time.Second)
        }
    })
    c.connection.AddCallback("PRIVMSG", func(e *ircenv.Event) {
		c.ReceivedMag(e)
		fmt.Println("PRIVMSG Channel:" + e.Arguments[0] + " " + e.Nick + ": " + e.Message())
	})
	c.connection.AddCallback("PUBMSG", func(e *ircenv.Event) {
		// TODO Not Yet
	})
}

func (c *ircClient) Join(channel string) {
    c.connection.Join(channel)
    fmt.Println("Joined" + channel)
}

func (c *ircClient) ReceivedMag(e *ircenv.Event) {
	go c.SaveMag(e)
}
func (c *ircClient) SaveMag(e *ircenv.Event) {
	msg := &Model.Message{}
	msg.Message = e.Message()
	msg.Channel = e.Arguments[0]
	msg.Nick = e.Nick
	msg.Time = time.Now().Local()
    msg.Server = c.server
	c.db.SaveMessage(msg)
}

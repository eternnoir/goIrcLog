package client

import (
	"fmt"
	ircenv "github.com/thoj/go-ircevent"
	DB "goIrcLog/Database"
    Model "goIrcLog/Model"
    "time"
)

type ircClient struct {
	channel, server, user, nickName string
	connection                      *ircenv.Connection
	db                              DB.DbProvider
}

func CreateClient(server, channel, user, nickName string,Db DB.DbProvider) *ircClient {
	ret := &ircClient{channel: channel, server: server, user: user, nickName: nickName}
	ret.connection = ircenv.IRC(ret.nickName, ret.user)
    ret.db = Db

	return ret
}

func (c *ircClient) Connect() bool {
	serverStr := c.server + ":6667"
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
		c.Join()
		fmt.Println("Joined" + c.channel)
	})
	c.connection.AddCallback("PRIVMSG", func(e *ircenv.Event) {
		c.ReceivedMag(e)
        fmt.Println("PRIVMSG Channel:"+ e.Arguments[0]+" " + e.Nick + ": " + e.Message())
	})
	c.connection.AddCallback("PUBMSG", func(e *ircenv.Event) {
		// TODO Not Yet
	})
}

func (c *ircClient) Join() {
	c.connection.Join(c.channel)
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
    c.db.SaveMessage(msg)
}

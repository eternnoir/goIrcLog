package client

import(
    config "goIrcLog/Config"
    DB "goIrcLog/Database"
	ircenv "github.com/thoj/go-ircevent"
)
func GetNewClient(c *config.Config) *ircClient{
    db := DB.GetNewDbProvider(c.DbType,c.DbConnection)
    ret := CreateClient(c.Server,c.Port,c.Channels,c.User,c.Nick,db) 
    return ret
}

func CreateClient(server,port, channel, user, nickName string,Db DB.DbProvider) *ircClient {
    ret := &ircClient{channel: channel, server: server, port:port,user: user, nickName: nickName}
	ret.connection = ircenv.IRC(ret.nickName, ret.user)
    ret.db = Db

	return ret
}



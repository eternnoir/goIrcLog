package client

import(
    config "goIrcLog/Config"
    DB "goIrcLog/Database"
	ircenv "github.com/thoj/go-ircevent"
    "strings"
)
func GetNewClient(c *config.Config) *ircClient{
    db := DB.GetNewDbProvider(c.DbType,c.DbConnection)
    chs := strings.Split(c.Channels,",")
    ret := CreateClient(c.Server,c.Port,chs,c.User,c.Nick,db) 
    return ret
}

func CreateClient(server string,port string, channel []string, user string, nickName string,Db DB.DbProvider) *ircClient {
    ret := &ircClient{channels: channel, server: server, port:port,user: user, nickName: nickName}
	ret.connection = ircenv.IRC(ret.nickName, ret.user)
    ret.db = Db

	return ret
}



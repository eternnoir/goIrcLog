goIrcLog
========

An Irc log bot, it can save channel massage to database(only support sqlite now). 

## Build

```
$ go get github.com/thoj/go-ircevent
$ go get code.google.com/p/go-sqlite/go1/sqlite3
$ go build 
```

## Config file

```json
{
    "Configs": [
    {
        "DbType": "Sqlite",
        "DbConnection": "./test.db",
        "Channels": "#goIrcLog,#goIrcLog2",
        "User": "logBotUser",
        "Password": "logBotPwd",
        "Nick": "logBotP",
        "Server": "chat.freenode.net",
        "Port": "6667"
    }
    ]
}
```

Channels split by ','


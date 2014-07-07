package main

import(
    "fmt"
    "./client"
	//"crypto/tls"
	//"testing"
	//"time"
)


func createIrcBot(channel string){
    c := client.CreateClient("chat.freenode.net","#","goIrclog","goIrcLog")
    c.Connect()
}

func main() {
    createIrcBot("HI")
}

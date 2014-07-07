package main

import(
    "fmt"
    "./client"
	//"crypto/tls"
	//"testing"
	//"time"
)


func createIrcBot(channel string){
    c := client.CreateClient("chat.freenode.net","#CS_RDSS","goIrclog","goIrcLog")
    c.Connect()
    fmt.Println("YA")
}



func main() {
    createIrcBot("HI")
}

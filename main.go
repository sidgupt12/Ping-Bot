package main

import (
	"fmt"

	"github.com/sidgupt12/Ping-Bot/bot"
	"github.com/sidgupt12/Ping-Bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return

}

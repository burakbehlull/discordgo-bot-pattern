package main

import (
	"botpattern/utils"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
)

var (
	client *discord.Session
)

func main() {
	utils.LoadDotenv()
	flag.Parse()
	
	token := os.Getenv("TOKEN")
	c, err := discord.New("Bot " + token)
	if err != nil {
		fmt.Println("Hata: ", err)
	}

	client = c

	err = client.Open()
	if err != nil {
		fmt.Println("Bot giriş yapmadı!")
	}

	fmt.Println("Bot giriş yaptı!")
	shutdownChannel()
}

func shutdownChannel() {
	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)

	<-sc
}

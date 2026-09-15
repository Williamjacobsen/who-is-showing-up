package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/williamjacobsen/who-is-showing-up/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env file:", err)
	}

	dg, err := discordgo.New("Bot " + os.Getenv("BOT_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln("Error creating discord session:", err)
	}

	dg.AddHandler(internal.HandleDiscordMessages)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	err = dg.Open()
	if err != nil {
		log.Fatalln("Error opening websocket connection:", err)
	}
	defer dg.Close()

	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

package main

import (
	"fmt"
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
		fmt.Println("Failed to load .env file:", err)
		return
	}

	dg, err := discordgo.New("Bot " + os.Getenv("BOT_DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating discord session:", err)
		return
	}

	dg.AddHandler(internal.HandleDiscordMessages)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	err = dg.Open()
	defer dg.Close()
	if err != nil {
		fmt.Println("Error opening websocket connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

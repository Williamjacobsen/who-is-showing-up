package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	stealth "github.com/jonfriesen/playwright-go-stealth"
	"github.com/williamjacobsen/who-is-showing-up/internal"

	"github.com/mxschmitt/playwright-go"
)

const (
	targetSite = "https://bot.sannysoft.com"
	userAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"
)

func scrape() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalln("Error starting playwright:", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalln("Could not create launh Chromium:", err)
	}
	defer browser.Close()

	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		UserAgent: playwright.String(userAgent),
	})
	if err != nil {
		log.Fatalln("Could not create page:", err)
	}
	defer page.Close()

	err = stealth.InjectWithOptions(page, stealth.Options{
		ChromeStealth: true,
	})
	if err != nil {
		log.Fatalln("Could not inject stealth script:", err)
	}

	_, err = page.Goto(targetSite, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	if err != nil {
		log.Fatalln("Could not go to page:", err)
	}

	fmt.Scanln()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env file:", err)
	}

	scrape()

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

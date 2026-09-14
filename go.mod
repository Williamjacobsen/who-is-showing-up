module github.com/williamjacobsen/who-is-showing-up

go 1.26.7

require (
	github.com/bwmarrin/discordgo v0.29.0
	github.com/joho/godotenv v1.5.1
	github.com/jonfriesen/playwright-go-stealth v0.0.3
	github.com/mxschmitt/playwright-go v0.6201.1
)

require (
	github.com/deckarep/golang-set/v2 v2.8.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
)

replace github.com/jonfriesen/playwright-go-stealth => ./playwright-go-stealth

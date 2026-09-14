package internal

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleDiscordMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("Message: %q\n", m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case strings.Contains(m.Content, "!help"):
		_, err := s.ChannelMessageSend(m.ChannelID, "Go here dumbass:\n https://github.com/Williamjacobsen/who-is-showing-up")
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

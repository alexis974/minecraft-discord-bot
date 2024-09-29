package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func sendMessage(session *discordgo.Session, channelID string, message string) error {
	// Send the provided message to the channel
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Printf("Error sending message to channel: %v", err)
		return err
	}

	log.Println("Message sent to channel!")
	return nil
}

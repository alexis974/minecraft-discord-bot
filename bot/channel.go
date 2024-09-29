package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func findChannelID(session *discordgo.Session) (string, error) {
	// Fetch all channels for the specified guild
	channels, err := session.GuildChannels(GuildID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Iterate through the channels to find one named "test"
	for _, channel := range channels {
		if channel.Name == ChannelName {
			return channel.ID, nil
		}
	}

	// If no matching channel is found, log an error and return
	log.Println("No channel named 'test' found")
	return "", nil
}

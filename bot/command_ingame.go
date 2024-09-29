package bot

import (
	"fmt"

	server "github.com/alexis974/minecraft-discord-bot/server"
	"github.com/bwmarrin/discordgo"
)

func inGame(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverInfo := server.GetServerInfo(MCServerHOST)

	if serverInfo.Online {
		var msg string
		msg += fmt.Sprintf("There is currently %d player(s) online\n", serverInfo.Players.Online)
		for index := 0; index < len(serverInfo.Players.List); index += 1 {
			msg += serverInfo.Players.List[index].Name + "\n"
		}

		discord.ChannelMessageSend(message.ChannelID, msg)
	} else {
		discord.ChannelMessageSend(message.ChannelID, "Server is offline")
	}
}

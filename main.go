package main

import (
	"log"
	"os"

	bot "github.com/alexis974/minecraft-discord-bot/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Could not load dotenv")
	}

	bot.BotToken = os.Getenv("DISCORD_BOT_TOKEN")
	bot.GuildID = os.Getenv("DISCORD_GUILD_ID")
	bot.ChannelName = os.Getenv("DISCORD_CHANNEL_NAME")
	bot.MCServerHOST = os.Getenv("MINECRAFT_SERVER_HOST")
	bot.Run()
}

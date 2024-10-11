# Minecraft Discord Bot

This is a Discord bot that notifies users on a Discord server whenever a player joins or leaves a Minecraft server.

## Purpose

I created this project as an excuse to learn and practice Go (Golang). The bot itself is straightforward.

## Features

- Sends notifications in a Discord channel when players join or leave a Minecraft server.
- Provides the `!ingame` command, which lists the players currently online on the Minecraft server.

## Configuration

To get the bot up and running, you'll need to configure the following env variables:

- **DISCORD_BOT_TOKEN**: The Discord bot token. You can get this by creating a bot on the [Discord Developer Portal](https://discord.com/developers/applications).
- **DISCORD_GUILD_ID**: The Discord server ID (Guild ID) where the bot will operate.
- **DISCORD_CHANNEL_NAME**: The name of the text channel where the bot will send notifications about players joining or leaving the Minecraft server.
- **MINECRAFT_SERVER_HOST**: The host name or IP address of the Minecraft server that the bot will monitor.

### Example Configuration

```env
DISCORD_BOT_TOKEN="your_discord_bot_token"
DISCORD_GUILD_ID="your_discord_server_id"
DISCORD_CHANNEL_NAME="general"
MINECRAFT_SERVER_HOST="minecraft.example.com"
```

## Docker Image

You can also run the bot using Docker. A pre-built Docker image is available: [alexis974/minecraft-discord-bot](https://hub.docker.com/r/alexis974/minecraft-discord-bot)

```bash
docker run -it \
    --env DISCORD_BOT_TOKEN="XXXX" \
    --env DISCORD_GUILD_ID="XXXX" \
    --env DISCORD_CHANNEL_NAME="XXXX" \
    --env MINECRAFT_SERVER_HOST="XXXX" \
    alexis974/minecraft-discord-bot:latest
```

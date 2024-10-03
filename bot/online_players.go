package bot

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
	"time"

	server "github.com/alexis974/minecraft-discord-bot/server"
	"github.com/bwmarrin/discordgo"
)

func watchOnlinePlayer(discord *discordgo.Session) {
	ticker := time.NewTicker(2 * time.Minute) // Set up a ticker for every 2 minutes
	defer ticker.Stop()                       // Ensure the ticker stops when the function ends

	channelID, err := findChannelID(discord)
	if err != nil {
		log.Println("Could not get channel ID to notify when player join server")
	}

	for {
		checkOnlinePlayer(discord, channelID)
		<-ticker.C // Wait for the next tick
	}
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func getCurrentPlayersInfo(filePath string) server.PlayersInfo {
	serverInfo := server.GetServerInfo(MCServerHOST)

	file, err := json.MarshalIndent(serverInfo.Players, "", " ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		log.Println(err)
	}

	return serverInfo.Players
}

func getOldPlayersInfo(filePath string) server.PlayersInfo {
	isFileExist := checkFileExists(filePath)

	if !isFileExist {
		log.Println("Init file info")
		return getCurrentPlayersInfo(filePath)
	}

	bytee, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}

	var playersInfo server.PlayersInfo

	err = json.Unmarshal(bytee, &playersInfo)
	if err != nil {
		log.Println(err)
	}

	return playersInfo
}

// Check if a player exists in the list by Uuid
func playerExists(player server.Player, players []server.Player) bool {
	for _, p := range players {
		if p.Uuid == player.Uuid {
			return true
		}
	}
	return false
}

// PlayerStatus holds the information of players who joined or left
type PlayerStatus struct {
	Joined []server.Player
	Left   []server.Player
}

func getPlayersStatus(oldPlayersInfo server.PlayersInfo, currentPlayersInfo server.PlayersInfo) PlayerStatus {
	var playerStatus PlayerStatus

	// Find players who joined
	for _, player := range currentPlayersInfo.List {
		if !playerExists(player, oldPlayersInfo.List) {
			playerStatus.Joined = append(playerStatus.Joined, player)
		}
	}

	// Find players who left
	for _, player := range oldPlayersInfo.List {
		if !playerExists(player, currentPlayersInfo.List) {
			playerStatus.Left = append(playerStatus.Left, player)
		}
	}

	return playerStatus
}

func checkOnlinePlayer(discord *discordgo.Session, channelID string) {

	var filePath string = "players.json"

	oldPlayersInfo := getOldPlayersInfo(filePath)
	currentPlayersInfo := getCurrentPlayersInfo(filePath)

	if reflect.DeepEqual(currentPlayersInfo.List, oldPlayersInfo.List) {
		log.Println("Nothing new on players list")
		return
	}

	playerStatus := getPlayersStatus(oldPlayersInfo, currentPlayersInfo)

	// Notify for every player that joined
	for _, player := range playerStatus.Joined {
		msg := player.Name + " joined the game !"
		log.Println(msg)
		sendMessage(discord, channelID, msg)
	}

	// Notify for every player that left
	for _, player := range playerStatus.Left {
		msg := player.Name + " has left the game !"
		log.Println(msg)
		sendMessage(discord, channelID, msg)
	}
}

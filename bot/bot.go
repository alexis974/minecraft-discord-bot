package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {

	// create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// add a event handler
	discord.AddHandler(newMessage)

	// open session
	discord.Open()
	defer discord.Close() // close session, after function termination

	go watchPlayer(discord)

	// keep bot running untill there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

// Function to send a message to a specific channel every two minutes
func sendMessageEveryTwoMinutes(discord *discordgo.Session, guildID, channelName, message string) {
	ticker := time.NewTicker(10 * time.Second) // Set up a ticker for every 2 minutes
	// ticker := time.NewTicker(2 * time.Minute) // Set up a ticker for every 2 minutes
	defer ticker.Stop() // Ensure the ticker stops when the function ends

	for {
		<-ticker.C // Wait for the next tick
		sendMessageToChannel(discord, guildID, "test", "Bonjour")
	}
}

func watchPlayer(discord *discordgo.Session) {
	guildID := ""

	sendMessageEveryTwoMinutes(discord, guildID, "test", "Bonjour")

	// sendMessageToChannel(discord, guildID, "test", "Bonjour")
}

func findChannelID(session *discordgo.Session, guildID string, channelName string) string {
	// Fetch all channels for the specified guild
	channels, err := session.GuildChannels(guildID)
	if err != nil {
		fmt.Println("111")
		// return err
	}

	// Iterate through the channels to find one named "test"
	for _, channel := range channels {
		if channel.Name == channelName {
			return channel.ID
		}
	}

	// If no matching channel is found, log an error and return
	log.Println("No channel named 'test' found")
	return "$$"
	// return nil
}

// Send a message to the channel
func sendMessageToChannel(session *discordgo.Session, guildID string, channelName string, message string) error {
	channelID := findChannelID(session, guildID, channelName)

	// Send the provided message to the channel
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Printf("Error sending message to channel: %v", err)
		return err
	}

	log.Println("Message sent to \"" + channelName + "\" channel!")
	return nil
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	/* prevent bot responding to its own message
	this is achived by looking into the message author id
	if message.author.id is same as bot.author.id then just return
	*/
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// respond to user message if it contains `!ingame`
	switch {
	case strings.Contains(message.Content, "!ingame"):
		serverInfo := getServerInfo("minecraft.alexisboissiere.fr")
		// watchPlayer(discord)

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
}

// A Response struct to map the Entire Response
type MCServerInfo struct {
	Online  bool        `json:"online"`
	Players PlayersInfo `json:"players"`
}

type PlayersInfo struct {
	Online int      `json:"online"`
	Max    int      `json:"max"`
	List   []Player `json:"list"`
}

type Player struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

func getServerInfo(serverName string) MCServerInfo {
	response, err := http.Get("http://api.mcsrvstat.us/3/" + serverName)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject MCServerInfo
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}


package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Bot token from Discord Developer Portal
var BotToken = "YOUR_BOT_TOKEN_HERE"

func main() {
	// Create a new Discord session using the bot token
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// Register the messageCreate function as a callback for the MessageCreate events
	dg.AddHandler(messageCreate)

	// Open a WebSocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening WebSocket connection,", err)
		return
	}
	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait here until CTRL+C or other termination signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Cleanly close down the Discord session
	dg.Close()
}

// This function will be called (as an event handler) every time a new message is created on any channel the bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Respond to the "!ping" command
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}

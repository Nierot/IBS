package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	Quote       string `json:"Quote"`
	AddedBy     string `json:"AddedBy"`
	MessageID   string `json:"MessageID"`
	Timestamp   int64  `json:"Timestamp"`
	Attachments string `json:"Attachments"`
	Reactions   string `json:"Reactions"`
}

type QuoteJSON struct {
	Quote       string   `json:"Quote"`
	AddedBy     string   `json:"AddedBy"`
	MessageID   string   `json:"MessageID"`
	Timestamp   int64    `json:"Timestamp"`
	Attachments []string `json:"Attachments"`
	Reactions   []string `json:"Reactions"`
}

func FindQuotes() {

}

func QuoteScanner(discord *discordgo.Session) {
	var (
		channel  = viper.GetString("Discord.QuoteChannel")
		interval = viper.GetInt("Discord.Scan.Interval")
	)
	fmt.Println("[DiscordScanner] Scanning for quotes on discord...")

	channelMsgs, err := discord.ChannelMessages(channel, 100, "", "", "")

	if err != nil {
		fmt.Println("[DiscordScanner] Error scanning channel!")
		panic(err)
	}

	for _, msg := range channelMsgs {
		var (
			attachments = make([]string, 0)
			reactions   = make([]string, 0)
		)

		for _, reaction := range msg.Reactions {
			reactions = append(reactions, reaction.Emoji.ID)
		}

		for _, attachment := range msg.Attachments {
			attachments = append(attachments, attachment.URL)
		}

		var quote = Quote{
			Quote:       msg.Content,
			AddedBy:     msg.Author.Username,
			MessageID:   msg.ID,
			Timestamp:   msg.Timestamp.Unix(),
			Attachments: sliceToString(attachments),
			Reactions:   sliceToString(reactions),
		}

		// create or insert
		if DB.Model(&Quote{}).Where("message_id = ?", msg.ID).Updates(&quote).RowsAffected == 0 {
			DB.Create(&quote)
		}

	}

	// restart this function after (interval) minutes
	time.AfterFunc(time.Duration(interval)*time.Minute, func() {
		QuoteScanner(discord)
	})
}

func sliceToString(s []string) string {
	arr := "["
	for _, str := range s {
		arr += str + ","
	}
	arr += "]"
	return arr
}

func StringToSlice(s string) []string {
	str := strings.Split(s, "[")[1]
	str = strings.Split(str, "]")[0]

	spl := strings.Split(str, ",")

	return spl[:len(spl)-1]
}

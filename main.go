package main

import (
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	hbot "github.com/otaviokr/hellivabot"
	log "github.com/sirupsen/logrus"

	"github.com/otaviokr/botaviokr-twitch-bot/trigger"
	"github.com/spf13/viper"
)

func main() {
	log.Warn("reading configuration file")
	readConfig()

	logFilename := viper.GetString("log.path")
	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.StandardLogger().Out = file
	} else {
		log.WithFields(
			log.Fields{
				"err": err.Error(),
				"file": logFilename,
			}).Info("failed to log to file, using default stderr")
	}
	defer file.Close()

	log.Warn("starting bot...")

	rawLevel := viper.GetString("log.level")
	logLevel, err := log.ParseLevel(rawLevel)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err.Error(),
				"raw_level": rawLevel,
			}).Fatal("failed to read configuration file")
	}
	log.SetLevel(logLevel)

	target := viper.GetString("irc.target")
	nickname := viper.GetString("irc.nickname")

	configurationOptions := func(bot *hbot.Bot) {
		bot.ThrottleDelay = 350 * time.Millisecond
		bot.Channels = viper.GetStringSlice("irc.channels")
		bot.HijackSession = !viper.GetBool("irc.ssl")
		bot.SSL = viper.GetBool("irc.ssl")
		password := viper.GetString("irc.password")
		if len(password) > 0 {
			bot.Password = password
		}
	}

	mybot, err := hbot.NewBot(target, nickname, configurationOptions)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err.Error(),
				"target": target,
				"nickname": nickname,
			}).Fatal("failed to connect")
	}

	owner := viper.GetString("triggers.bot.owner")
	repo := viper.GetString("triggers.bot.repository")
	mybot.AddTrigger(trigger.Bot(owner, repo))

	streamHolicsFriends := viper.GetStringSlice("triggers.streamholics.friends")
	mybot.AddTrigger(trigger.StreamHolicsJoin(streamHolicsFriends))

	mybot.AddTrigger(trigger.SayHello())

	mybot.Run()
}

// ReadConfig will parse the properties file.
func readConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("botaviokr-twitch-bot")
	viper.SetConfigType("yaml")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.WithFields(
			log.Fields{
				"file": e.Name,
				"event": e.Op.String(),
			}).Info("configuration file changed and settings have been refreshed")
	})

	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err.Error(),
			}).Fatal("failed to process config file")
	}
}

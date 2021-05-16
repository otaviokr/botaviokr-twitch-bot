# botaviokr-twitch-bot
A IRC chat bot for Twitch chat, written in Go

# Overview

This is my Twitch bot to help me and the viewers during livestreams. FAQ, links, some automated process and much more can be delegated to the bot and it will handle it while I can focus in on the live stream.

It is actually my extended, customized version of another bot I created, the [hellivabot](https://github.com/otaviokr/hellivabot), which is itself an customized version of the [Hellabot](https://github.com/whyrusleeping/hellabot), an excellent, more generic IRC chat bot. So if you are looking for a generic IRC bot, check Hellabot; if you want you own Twitch chat bot, check hellivabot; if you want some ready to go, feel free to fork the BOTavio_KR, this project.

# How to run it

## Stand-alone program

Being just an extended version of **hellivabot**, you should be able to run this one the same way. However, we need more properties in the configuration file, to make the triggers work as expected. Check out the YAML template to see which properties are expected.

You can run it as stand-alone program, but you need to have [Go](https://golang.org) installed and configured in your system. Compiling is optional, so the quickest way to have it running is just:

```bash
[user@host botaviokr-twitch-bot]$ go run main.go
```

## Containerized

You can run it in a container. The main advantage is that you don't have to worry about anything other than having [Docker](https://docker.com) or [Podman](https://podman.org) installed and configured. The host ports defined below are just a suggestion - feel free to adapt them to your system needs. If you have podman installed, just change the command (call podman instead of docker) and everything should work the same.

```bash
[user@host botaviokr-twitch-bot]$ docker build .
[user@host botaviokr-twitch-bot]$ docker run -d --name botaviokr -p  -v $(pwd)/logs:/botaviokr/logs -v $(pwd)/config:/botaviokr/config
```

# Triggers

If you want more details on how to implement new functionalities for the chat bot, check the README in the hellivabot. If you are new here, a trigger is how the bot knows when to take action, and what to do. A trigger has two components:

- the **condition** is a test that determines if the bot should act or not. For example, when a specific message is written in the chat, or a user joined, or if the number of people in the chat is a new record, it returns `true` and the action will be performed;

- the **action** is the commands that the bot will perform if the condition was true. Examples of actions are the bot writing in the chat, saving some data in a file, updating a variable in the program, running a query in a database etc.

My bot has the triggers I personally find interesting - if you find some of them useless or unnecessary, that's OK and it is very easy to deactivated them (they are added in the `main.go` file). If you miss a feature, open an issue and let me know - I love new ideas.

# Auxiliary Technologies

## MQTT - Mosquitto

All data collected by the bot that I would like to pass it on to other systems (e.g. databases), are sent to a MQTT. More specifically, I have a [Mosquitto](https://mosquitto.org) instance running on my network, but the bot itself can connect to any MQTT broker.

The MQTT client is instantiated in the `main.go` and is used in the triggers, such as in `guestbook.go`.

The MQTT details are expected to be defined in the YAML properties file. If those details are missing, the bot will have the trigger disabled.

My opinion is that the bot should be focused on the chat, so there is not listener to consume messages from MQTT, just to publish to it.
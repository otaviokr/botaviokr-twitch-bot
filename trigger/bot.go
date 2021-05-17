package trigger

import (
	"fmt"
	"strings"

	hbot "github.com/otaviokr/hellivabot"
	"github.com/spf13/viper"
)

// Bot will respond some details about the bot
func Bot() hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!bot")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(
				m, fmt.Sprintf(
					"Oi, eu sou o Bot do %s. Se quiser realmente me conhecer, o meu código está disponível em %s",
					viper.GetString("triggers.bot.owner"),
					viper.GetString("triggers.bot.repository")))
			return false
		},
	}
}

// Commands will present all the commands the bot responds to.
func Commands() hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!commands")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, "Estes são os comandos que eu entendo:\n!bot - meus detalhes\n!hi - manda um oi\n!hello - manda um olá\n!twitter - me acha lá\n!github - meus repositórios")
			return true
		},
	}
}

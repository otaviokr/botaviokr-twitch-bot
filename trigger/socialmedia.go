package trigger

import (
	"fmt"
	"strings"

	hbot "github.com/otaviokr/hellivabot"
	"github.com/spf13/viper"
)

// Github will respond some details about the bot
func Github() hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!github")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m,
				fmt.Sprintf(
					"Programas, scripts, bots e outros códigos que eu faço aqui estão disponíveis no meu Github: %s",
					viper.GetString("triggers.socialmedia.github")))
			return true
		},
	}
}

// Twitter will give the link to Twitter profile.
func Twitter() hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!twitter")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m,
				fmt.Sprintf(
					"Quer conversar fora da Live? Me chame lá no Twitter: %s",
					viper.GetString("triggers.socialmedia.twitter")))
			return true
		},
	}
}

// Youtube will give the link to the youtube channel.
func Youtube() hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!youtube")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m,
				fmt.Sprintf(
					"O conteúdo desta live e muito mais está disponível no canal do Youtube: %s",
					viper.GetString("triggers.socialmedia.youtube")))
			return true
		},
	}
}

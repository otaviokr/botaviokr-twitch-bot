package trigger

import (
	"fmt"
	"strings"

	hbot "github.com/otaviokr/hellivabot"
)

// Github will respond some details about the bot
func Github(github string) hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!github")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, fmt.Sprintf("Programas, scripts, bots e outros códigos que eu faço aqui estão disponíveis no meu Github: %s", github))
			return true
		},
	}
}

// Twitter will give the link to Twitter profile.
func Twitter(twitter string) hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!twitter")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, fmt.Sprintf("Quer conversar fora da Live? Me chame lá no Twitter: %s", twitter))
			return true
		},
	}
}

// Youtube will give the link to the youtube channel.
func Youtube(youtube string) hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!youtube")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, fmt.Sprintf("O conteúdo desta live e muito mais está disponível no canal do Youtube: %s", youtube))
			return true
		},
	}
}

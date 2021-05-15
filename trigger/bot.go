package trigger

import (
	"fmt"
	"strings"

	hbot "github.com/otaviokr/hellivabot"
)

// Bot will respond some details about the bot
func Bot(owner, repo string) hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && strings.EqualFold(strings.TrimSpace(m.Content), "!bot")
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, fmt.Sprintf("Oi, eu sou o Bot do %s. Se quiser realmente me conhecer, o meu código está disponível em %s", owner, repo))
			return false
		},
	}
}
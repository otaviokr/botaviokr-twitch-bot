package trigger

import (
	"fmt"
	"regexp"
	"strings"

	hbot "github.com/otaviokr/hellivabot"
)

// SayHello is a very simple trigger that answers when someone says hi with "!hi" or "!hello"
func SayHello() hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			return strings.EqualFold(m.Command, "PRIVMSG") && regexp.MustCompile("!h(i|ello)").MatchString(m.Content)
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, fmt.Sprintf("Oh, hi there %s. Nice to see you here!", m.From))
			return false
		},
	}
}

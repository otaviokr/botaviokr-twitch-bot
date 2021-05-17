package trigger

import (
	"fmt"
	"strings"

	hbot "github.com/otaviokr/hellivabot"
	"github.com/spf13/viper"
)

// StreamHolicsJoin will shout whenever a friend from StreamHolics joins the chat.
func StreamHolicsJoin() hbot.Trigger {
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			if !strings.EqualFold(m.Command, "JOIN") {
				return false
			}

			friendList := viper.GetStringSlice("triggers.streamholics.friends")
			for _, friend := range friendList {
				if strings.EqualFold(strings.TrimSpace(m.Content), friend) {
					return true
				}
			}
			return false
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, fmt.Sprintf("!sh %s", m.Content))
			return false
		},
	}
}
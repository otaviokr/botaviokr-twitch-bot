package trigger

import (
	"fmt"
	"strings"

	hbot "github.com/otaviokr/hellivabot"
)

// StreamHolicsJoin will shout whenever a friend from StreamHolics joins the chat.
func StreamHolicsJoin(friendList []string) hbot.Trigger {
	friendJoined := ""
	return hbot.Trigger{
		Condition: func (b *hbot.Bot, m *hbot.Message) bool {
			if !strings.EqualFold(m.Command, "JOIN") {
				return false
			}

			for _, friend := range friendList {
				if strings.EqualFold(strings.TrimSpace(m.Content), friend) {
					friendJoined = friend
					return true
				}
			}
			return false
		},
		Action: func (b *hbot.Bot, m *hbot.Message) bool {
			b.Reply(m, fmt.Sprintf("!sh %s", friendJoined))
			return false
		},
	}
}
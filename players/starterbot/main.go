// Bruno's Planet Wars bot

package main

import (
	"github.com/brvoisin/planetwarsbot"
)

func main() {
	bot := NewStarterBot()
	planetwarsbot.Run(bot)
}

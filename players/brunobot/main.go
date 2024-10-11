// Bruno's Planet Wars bot

package main

import (
	"github.com/brvoisin/planetwars"
)

func main() {
	bot := NewBrunoBot()
	planetwars.Run(bot)
}

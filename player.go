package planetwarsbot

import "os"

type Player interface {
	DoTurn(Map) []Order
}

// Run the bot forever.
func Run(player Player) {
	for {
		planetMap := ParseInputMap(os.Stdin)
		orders := player.DoTurn(planetMap)
		SerializeOrders(orders, os.Stdout)
	}
}

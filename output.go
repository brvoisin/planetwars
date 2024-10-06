package planetwars

import (
	"fmt"
	"io"
)

func SerializeOrders(orders []Order, writer io.Writer) {
	for _, order := range orders {
		fmt.Fprintf(writer, "%d %d %d\n", order.Source, order.Dest, order.Ships)
	}
	fmt.Fprintln(writer, "go")
}

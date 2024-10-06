package planetwars_test

import (
	"bytes"
	"testing"

	"github.com/brvoisin/planetwars"
)

func TestSerializeOrders(t *testing.T) {
	type args struct {
		orders []planetwars.Order
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no order",
			args: args{orders: []planetwars.Order{}},
			want: "go\n",
		},
		{
			name: "specification example",
			args: args{orders: []planetwars.Order{
				{Source: 1, Dest: 17, Ships: 50},
				{Source: 4, Dest: 17, Ships: 50},
			}},
			want: "1 17 50\n4 17 50\ngo\n",
		},
		{
			name: "nil",
			args: args{orders: nil},
			want: "go\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			planetwars.SerializeOrders(tt.args.orders, writer)
			if gotWriter := writer.String(); gotWriter != tt.want {
				t.Errorf("serializeOrders() = %#v, want %#v", gotWriter, tt.want)
			}
		})
	}
}

package planetwarsbot_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/brvoisin/planetwarsbot"
)

const input = `P 0    0    1 34 2
P 7    9    2 34 2
P 3.14 2.71 0 15 5

F 1 15 0 1 12 2
F 2 28 1 2  8 4
go
`

func TestParseInputMap(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want planetwarsbot.Map
	}{
		{
			name: "specification example",
			args: args{input: strings.NewReader(input)},
			want: expectedMap(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := planetwarsbot.ParseInputMap(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInputMap() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func expectedMap() planetwarsbot.Map {
	return planetwarsbot.Map{
		Planets: []planetwarsbot.Planet{
			{
				ID:       0,
				Position: planetwarsbot.Point{X: 0, Y: 0},
				Owner:    1,
				Ships:    34,
				Growth:   2,
			},
			{
				ID:       1,
				Position: planetwarsbot.Point{X: 7, Y: 9},
				Owner:    2,
				Ships:    34,
				Growth:   2,
			},
			{
				ID: 2,
				Position: planetwarsbot.Point{
					X: 3.140000104904175,
					Y: 2.7100000381469727,
				}, // Precision issue with float64 (not float32)
				Owner:  0,
				Ships:  15,
				Growth: 5,
			},
		},
		Fleets: []planetwarsbot.Fleet{
			{
				Owner:         1,
				Ships:         15,
				Source:        0,
				Dest:          1,
				TotalTurn:     12,
				RemainingTurn: 2,
			},
			{
				Owner:         2,
				Ships:         28,
				Source:        1,
				Dest:          2,
				TotalTurn:     8,
				RemainingTurn: 4,
			},
		},
	}
}

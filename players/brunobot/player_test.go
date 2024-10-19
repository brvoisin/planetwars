package main

import (
	"reflect"
	"testing"

	"github.com/brvoisin/planetwars"
)

func TestDoTurn(t *testing.T) {
	type args struct {
		planetMap planetwars.Map
	}
	tests := []struct {
		name string
		args args
		want []planetwars.Order
	}{
		{
			name: "prefer nearest planet",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 6, Y: 0},
						Owner:    planetwars.Neutral,
						Ships:    50,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Neutral,
						Ships:    50,
						Growth:   1,
					},
				},
			}},
			want: []planetwars.Order{
				{Source: 0, Dest: 2, Ships: 51},
			},
		},
		{
			name: "prefer planet with bigger growth",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 6, Y: 0},
						Owner:    planetwars.Neutral,
						Ships:    50,
						Growth:   5,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Neutral,
						Ships:    50,
						Growth:   1,
					},
				},
			}},
			want: []planetwars.Order{
				{Source: 0, Dest: 1, Ships: 51},
			},
		},
		{
			name: "multiple orders",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    80,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 2, Y: 0},
						Owner:    planetwars.Neutral,
						Ships:    30,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 5, Y: 5},
						Owner:    planetwars.Neutral,
						Ships:    30,
						Growth:   1,
					},
					{
						ID:       3,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Neutral,
						Ships:    30,
						Growth:   1,
					},
				},
			}},
			want: []planetwars.Order{
				{Source: 0, Dest: 1, Ships: 31},
				{Source: 0, Dest: 3, Ships: 31},
			},
		},
		{
			name: "consider growth",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Opponent,
						Ships:    50,
						Growth:   2,
					},
				},
			}},
			want: []planetwars.Order{
				// The ships on opponent planet will increase by 10 in 5 turns.
				{Source: 0, Dest: 1, Ships: 61},
			},
		},
		{
			name: "consider fleets arrived before my new fleet",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Opponent,
						Ships:    50,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 10, Y: 10},
						Owner:    planetwars.Opponent,
						Ships:    100,
						Growth:   1,
					},
				},
				Fleets: []planetwars.Fleet{
					{
						Owner:         planetwars.Myself,
						Ships:         10,
						Source:        0,
						Dest:          1,
						TotalTurn:     5,
						RemainingTurn: 2,
					},
					{
						Owner:         planetwars.Opponent,
						Ships:         15,
						Source:        2,
						Dest:          1,
						TotalTurn:     12,
						RemainingTurn: 3,
					},
					// The fleet below should be ignored.
					{
						Owner:         planetwars.Opponent,
						Ships:         15,
						Source:        2,
						Dest:          0,
						TotalTurn:     15,
						RemainingTurn: 3,
					},
				},
			}},
			want: []planetwars.Order{
				// Planet 1 ships
				// Current turn: 50
				// Turn 1: growth +1 = 51
				// Turn 2: growth +1 and fleet 0 -10 = 42
				// Turn 3: growth +1 and fleet 1 +15 = 58
				// Turn 4: growth +1 = 59
				// Turn 5: growth +1 = 60
				{Source: 0, Dest: 1, Ships: 61},
			},
		},
		{
			name: "don't send an order with zero ship",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Neutral,
						Ships:    20,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 10, Y: 10},
						Owner:    planetwars.Neutral,
						Ships:    30,
						Growth:   1,
					},
				},
				Fleets: []planetwars.Fleet{
					{
						Owner:         planetwars.Myself,
						Ships:         21,
						Source:        0,
						Dest:          1,
						TotalTurn:     5,
						RemainingTurn: 2,
					},
				},
			}},
			want: []planetwars.Order{
				// Don't send a fleet to the planet 1.
				{Source: 0, Dest: 2, Ships: 31},
			},
		},
		{
			name: "do not endanger a planet",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Neutral,
						Ships:    41,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 10, Y: 10},
						Owner:    planetwars.Neutral,
						Ships:    100,
						Growth:   1,
					},
				},
				Fleets: []planetwars.Fleet{
					{
						Owner:         planetwars.Opponent,
						Ships:         60,
						Source:        2,
						Dest:          0,
						TotalTurn:     5,
						RemainingTurn: 2,
					},
				},
			}},
			want: []planetwars.Order{},
		},
		{
			name: "do not repeat orders from multiple planets",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 10, Y: 10},
						Owner:    planetwars.Neutral,
						Ships:    50,
						Growth:   1,
					},
				},
			}},
			want: []planetwars.Order{
				{Source: 1, Dest: 2, Ships: 51},
			},
		},
		{
			name: "correctly consider growth for fleet ships",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Opponent,
						Ships:    50,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 0, Y: 7},
						Owner:    planetwars.Myself,
						Ships:    50,
						Growth:   1,
					},
				},
				Fleets: []planetwars.Fleet{
					{
						Owner:         planetwars.Myself,
						Ships:         50 + 1*2 + 1,
						Source:        2,
						Dest:          1,
						TotalTurn:     2,
						RemainingTurn: 1,
					},
				},
			}},
			want: []planetwars.Order{},
		},
		{
			name: "don't send fleet if the planet will be mine later",
			args: args{planetMap: planetwars.Map{
				Planets: []planetwars.Planet{
					{
						ID:       0,
						Position: planetwars.Point{X: 0, Y: 0},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       1,
						Position: planetwars.Point{X: 0, Y: 2},
						Owner:    planetwars.Myself,
						Ships:    100,
						Growth:   1,
					},
					{
						ID:       2,
						Position: planetwars.Point{X: 0, Y: 5},
						Owner:    planetwars.Opponent,
						Ships:    50,
						Growth:   1,
					},
				},
				Fleets: []planetwars.Fleet{
					{
						Owner:         planetwars.Myself,
						Ships:         50 + 1*5 + 1,
						Source:        0,
						Dest:          2,
						TotalTurn:     5,
						RemainingTurn: 4,
					},
				},
			}},
			want: []planetwars.Order{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := NewBrunoBot()
			if got := bot.DoTurn(tt.args.planetMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doTurn() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

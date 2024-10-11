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
			name: "multiple orders",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := NewBrunoBot()
			if got := bot.DoTurn(tt.args.planetMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

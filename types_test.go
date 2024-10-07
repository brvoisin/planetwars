package planetwars_test

import (
	"reflect"
	"testing"

	"github.com/brvoisin/planetwars"
)

func TestMapMyPlanets(t *testing.T) {
	type fields struct {
		Planets []planetwars.Planet
	}
	tests := []struct {
		name   string
		fields fields
		want   []planetwars.Planet
	}{
		{
			name: "filter my planets",
			fields: fields{
				Planets: []planetwars.Planet{
					{ID: 0, Owner: planetwars.Neutral},
					{ID: 1, Owner: planetwars.Myself},
					{ID: 2, Owner: planetwars.Opponent},
					{ID: 3, Owner: planetwars.Myself},
				},
			},
			want: []planetwars.Planet{
				{ID: 1, Owner: planetwars.Myself},
				{ID: 3, Owner: planetwars.Myself},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := planetwars.Map{
				Planets: tt.fields.Planets,
			}
			if got := m.MyPlanets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map.MyPlanets() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestMapMyFleets(t *testing.T) {
	type fields struct {
		Planets []planetwars.Planet
		Fleets  []planetwars.Fleet
	}
	tests := []struct {
		name   string
		fields fields
		want   []planetwars.Fleet
	}{
		{
			name: "filter my fleets",
			fields: fields{
				Fleets: []planetwars.Fleet{
					{Owner: planetwars.Opponent, Ships: 10},
					{Owner: planetwars.Myself, Ships: 20},
					{Owner: planetwars.Myself, Ships: 30},
				},
			},
			want: []planetwars.Fleet{
				{Owner: planetwars.Myself, Ships: 20},
				{Owner: planetwars.Myself, Ships: 30},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := planetwars.Map{
				Planets: tt.fields.Planets,
				Fleets:  tt.fields.Fleets,
			}
			if got := m.MyFleets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map.MyFleets() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestMapFleetsTo(t *testing.T) {
	type fields struct {
		Planets []planetwars.Planet
		Fleets  []planetwars.Fleet
	}
	type args struct {
		ID planetwars.PlanetID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []planetwars.Fleet
	}{
		{
			name: "filter fleet destination",
			fields: fields{
				Planets: []planetwars.Planet{
					{ID: 0, Owner: planetwars.Myself},
					{ID: 1, Owner: planetwars.Opponent},
				},
				Fleets: []planetwars.Fleet{
					{Source: 0, Dest: 1, Ships: 10},
					{Source: 1, Dest: 0, Ships: 20},
				},
			},
			args: args{ID: 0},
			want: []planetwars.Fleet{
				{Source: 1, Dest: 0, Ships: 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := planetwars.Map{
				Planets: tt.fields.Planets,
				Fleets:  tt.fields.Fleets,
			}
			if got := m.FleetsTo(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map.FleetsTo() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

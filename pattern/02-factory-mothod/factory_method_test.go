package factory_mothod

import (
	"reflect"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	type args struct {
		gameType GameType
	}
	tests := []struct {
		name string
		args args
		want *Game
	}{
		{
			name: "testcase 1",
			args: args{
				gameType: FpsGame,
			},
			want: &Game{
				Name: "沙滩排球",
			},
		},
		{
			name: "testcase 2",
			args: args{
				gameType: ActionGame,
			},
			want: &Game{
				Name: "死或生",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGame(tt.args.gameType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
